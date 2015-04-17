define(['util', 'i18n', 'can', 'text!app/preload/badwords.json'],
    function (util, i18n, can, badwords) {
    /**
     * A set for general purpose utilities.
     * @author dorajistyle
     * @module util
     */
    var badWords = $.parseJSON(badwords);
     var NUMERIC_REGEX = /^[0-9]+$/,
        INTEGER_REGEX = /^\-?[0-9]+$/,
        DECIMAL_REGEX = /^\-?[0-9]*\.?[0-9]+$/,
        USER_NAME_REGEX = /^[a-zA-Z0-9-.]+$/i,
        // EMAIL_REGEX = /^[a-zA-Z0-9.!#$%&amp;'*+\-\/=?\^_`{|}~\-]+@[a-zA-Z0-9\-]+(?:\.[a-zA-Z0-9\-]+)*$/,
        EMAIL_REGEX = /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|”(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*”)@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/,
        ALPHA_REGEX = /^[a-z]+$/i,
        ALPHA_NUMERIC_REGEX = /^[a-z0-9]+$/i,
        ALPHA_DASH_REGEX = /^[a-z0-9_\-]+$/i,
        NATURAL_REGEX = /^[0-9]+$/i,
        NATURAL_NO_ZERO_REGEX = /^[1-9][0-9]*$/i,
        IP_REGEX = /^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})$/i,
        BASE64_REGEX = /[^a-zA-Z0-9\/\+=]/i,
        NUMERIC_DASH_REGEX = /^[\d\-\s]+$/,
        BAD_WORD_REGEX = new RegExp(badWords.en.join('|'), 'gi'),
        BAD_WORD_INT_REGEX = new RegExp(badWords.en.join('|')+badWords.ko.join('|'), 'gi'),
        URL_REGEX = /^((http|https):\/\/(\w+:{0,1}\w*@)?(\S+)|)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$/;
    return {
        /**
         * Validate a filed that is checked.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isChecked: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var result = $field.is(':checked');
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },
        /**
         * Validate a email field that is correct.
         * @param name
         * @returns {boolean}
         */
        validateEmail: function (name, skipFocus) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = EMAIL_REGEX.test(value);
            util.logDebug('email value', value);
            util.logDebug('email result', result);
            if (!result) {
                util.putMessage(i18n.t('validation.email'));
                if (!skipFocus) $field.focus();
                return false;
            }
            return true;
        },
        /**
        * Validate a username field that is correct.
        * @param name
        * @returns {boolean}
        */
        validateUsernameOrEmail: function (name) {
          var $field = $('input[name=' + name + ']');
          var value = $field.val();
          var result = false;
          var emailResult = EMAIL_REGEX.test(value);
          if (!emailResult) {
            util.putMessage(i18n.t('validation.emailOr'));
          }
          result = (emailResult || this.validateUsername(name));
          if(result) {
              util.clearMessages();
          }
          return result;
        },

        /**
        * Validate a username field that is correct.
        * @param name
        * @returns {boolean}
        */
        validateUsername: function (name) {
          var $field = $('input[name=' + name + ']');
          var value = $field.val();
          var result = false;
          result = this.minLength(name, 4, 'validation.username', true) && this.maxLength(name, 16, 'validation.username', true) && this.isCorrectUserName(name, 'validation.username', true) && this.hasGoodEnglishWords(name, 'validation.badUsername', true);
          // if(result) {
            // util.clearMessages();
          // }
          return result;
        },
        /**
        * Validate a field's value that is a correct user name.
        * @param name
        * @param msg
        * @returns {boolean}
        */
        isCorrectUserName: function (name, msg, skipFocus) {
          var $field = $('input[name=' + name + ']');
          var value = $field.val();
          var result = USER_NAME_REGEX.test(value);
          if (!result) {
            util.putMessage(i18n.t(msg));
            if (!skipFocus) {
              $field.focus();
            }
            return false;
          }
          return true;
        },

        /**
        * Validate a field's length that is above the minimum Length.
        * @param name
        * @param length
        * @param msg
        * @returns {boolean}
        */
        minLength: function (name, length, msg, skipFocus) {
          var result = this.minLengthWithSpace(name, length, msg, skipFocus, false);
          return result;
        },

        /**
        * Validate a field's length that is above the minimum Length and if allowSpace parameter is false, it will ignore space as length.
        * @param name
        * @param length
        * @param msg
        * @param skipFocus
        * @param allowSpace
        * @returns {boolean}
        */
        minLengthWithSpace: function (name, length, msg, skipFocus, allowSpace) {
          if (!NUMERIC_REGEX.test(length)) {
            return false;
          }
          var $field = $('input[name=' + name + ']');
          var value = $field.val();
          var valueLength = $.trim(value).length;
          if (allowSpace) {valueLength = value.length; }
          var result = (valueLength >= parseInt(length, 10));
          if (!result) {
            util.putMessage(i18n.t(msg));
            if (!skipFocus) $field.focus();
            return false;
          }
          return true;
        },

        /**
         * Validate a textarea's length that is above the minimum Length.
         * @param name
         * @param length
         * @param msg
         * @returns {boolean}
         */
        minLengthTextarea: function (name, length, msg) {
            if (!NUMERIC_REGEX.test(length)) {
                return false;
            }
            var $textarea = $('textarea[name=' + name + ']');
            var value = $textarea.val();
            var result = (value.length >= parseInt(length, 10));
            if (!result) {
                util.putMessage(i18n.t(msg));
                $textarea.focus();
                return false;
            }
            return true;
        },

         /**
         * Validate a textarea of form's length that is above the minimum Length.
         * @param form
         * @param name
         * @param length
         * @param msg
         * @returns {boolean}
         */
        minLengthTextareaOfForm: function (form, name, length, msg) {
            if (!NUMERIC_REGEX.test(length)) {
                return false;
            }
            var $textarea = $(form.find('textarea[name=' + name + ']'));
            var value = $textarea.val();
            var result = ($.trim(value).length >= parseInt(length, 10));
            if (!result) {
                util.putMessage(i18n.t(msg));
                $textarea.focus();
                return false;
            }
            return true;
        },


        /**
         * Validate a field's length that is below the maximum Length.
         * @param name
         * @param length
         * @param msg
         * @returns {boolean}
         */
        maxLength: function (name, length, msg, skipFocus) {
            if (!NUMERIC_REGEX.test(length)) {
                return false;
            }
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = (value.length <= parseInt(length, 10));
            if (!result) {
                util.putMessage(i18n.t(msg));
                if (!skipFocus) {
                  $field.focus();
                }
                return false;
            }
            return true;
        },

        /**
        * validate identical
        * @param $target
        * @param length
        * @returns {boolean}
        */
        validateIdentical: function (left, right, msg) {
          var $left = $('input[name=' + left + ']');
          var $right = $('input[name=' + right + ']');
          var result = this.checkIdentical($left, $right);
          if (!result) {
            util.putMessage(i18n.t(msg));
          }
          return result;
        },
        /**
        * Check identical
        * @param $target
        * @param length
        * @returns {boolean}
        */
        checkIdentical: function ($left ,$right) {
          var leftValue = $left.val();
          var rightValue = $right.val();
          var result = (leftValue == rightValue);
          return result;
        },

        /**
        * Validate a field's length that is exactly same as length.
        * @param name
        * @param length
        * @param msg
        * @returns {boolean}
        */
        exactLength: function (name, length, msg) {
          var $field = $('input[name=' + name + ']');
          var result = this.checkExactLength($field, length);
          if (!result) {
            this.putMessage(i18n.t(msg));
            $field.focus();
            return false;
          }
          return true;
        },

        /**
        * Check exact length
        * @param $target
        * @param length
        * @returns {boolean}
        */
        checkExactLength: function ($target , length) {
          if (!NUMERIC_REGEX.test(length)) {
            return false;
          }
          var value = $target.val();
          var result = (value.length == parseInt(length, 10));
          return result;
        },


        /**
         * Validate a field's value that is grater than param value.
         * @param name
         * @param param
         * @param msg
         * @returns {boolean}
         */
        greaterThan: function (name, param, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = (DECIMAL_REGEX.test(value) && (parseFloat(value) > parseFloat(param)));
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is less than param value.
         * @param name
         * @param param
         * @param msg
         * @returns {boolean}
         */
        lessThan: function (name, param, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = (DECIMAL_REGEX.test(value) && (parseFloat(value) < parseFloat(param)));
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is alphabet.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isAlpha: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = ALPHA_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is alphabet or Numeric.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isAlphaNumeric: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = ALPHA_NUMERIC_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is alphabet or underbar or dash.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isAlphaDash: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = ALPHA_DASH_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },
        /**
         * Validate a field's value that is numeric.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isNumeric: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = DECIMAL_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is integer.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isInteger: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = INTEGER_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is decimal.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isDecimal: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = DECIMAL_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is natural.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isNatural: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = NATURAL_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is natural without zero.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        isNaturalNoZero: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = NATURAL_NO_ZERO_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is correct ip address.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        validateIp: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = IP_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is correct Base64 format.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        validateBase64: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = BASE64_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is correct Url.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        validateUrl: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var result = URL_REGEX.test(value);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },

        /**
         * Validate a field's value that is correct credit card number.
         * @param name
         * @param msg
         * @returns {boolean}
         */
        validateCreditCard: function (name, msg) {
            var $field = $('input[name=' + name + ']');
            var value = $field.val();
            var nCheck = 0, nDigit = 0, bEven = false;
            var strippedField = value.replace(/\D/g, '');
            for (var n = strippedField.length - 1; n >= 0; n--) {
                var cDigit = strippedField.charAt(n);
                nDigit = parseInt(cDigit, 10);
                if (bEven) {
                    if ((nDigit *= 2) > 9) nDigit -= 9;
                }
                nCheck += nDigit;
                bEven = !bEven;
            }
            var result = (NUMERIC_DASH_REGEX.test(value) && ((nCheck % 10) === 0));
            if (!result) {
                util.putMessage(i18n.t(msg));
                $field.focus();
                return false;
            }
            return true;
        },
        /**
         * Validate two field's value there are identical.
         * @param nameSrc
         * @param nameDest
         * @param msg
         * @returns {boolean}
         */
        isIdentical: function (nameSrc, nameDest, msg) {
            var $fieldSrc = $('input[name=' + nameSrc + ']');
            var valueSrc = $fieldSrc.val();
            var $fieldDest = $('input[name=' + nameDest + ']');
            var valueDest = $fieldDest.val();
            var result = (valueSrc == valueDest);
            if (!result) {
                util.putMessage(i18n.t(msg));
                $fieldDest.focus();
                return false;
            }
            return true;
        },



                /**
         * check field has good words
         * @param name
         * @param msg
         * @returns {boolean}
         */
        hasGoodEnglishWords: function (name, msg, skipFocus) {
            var $field = $('input[name=' + name + ']');
            var value = $.trim($field.val());
            var resultTest = BAD_WORD_REGEX.test(value);
            var resultExec = BAD_WORD_REGEX.exec(value);
//            util.logDebug('regex', BAD_WORD_REGEX);
//            util.logDebug('value', value);
//            util.logDebug('result words', resultTest);
//            util.logDebug('exec length', resultExec);

            if (resultTest || resultExec != undefined) {
                util.putMessage(i18n.t(msg));
                if (!skipFocus) { $field.focus(); }
                return false;
            }
            return true;
        },
        hasGoodWordsTextArea: function (name, msg, skipFocus) {
            var $field = $('textarea[name=' + name + ']');
            var value = $.trim($field.val());
            var resultTest = BAD_WORD_REGEX.test(value);
            var resultExec = BAD_WORD_REGEX.exec(value);
//            util.logDebug('regex', BAD_WORD_REGEX);
//            util.logDebug('value', value);
//            util.logDebug('result words', resultTest);
//            util.logDebug('exec length', resultExec);

            if (resultTest || resultExec != undefined) {
                util.putMessage(i18n.t(msg));
                if (!skipFocus) { $field.focus(); }
                return false;
            }
            return true;
        },
        /**
         * replace bad words to stars
         * @param str
         * @returns {*|XML|string|void}
         */
        replaceBadWords: function (str) {
            return str.replace(BAD_WORD_INT_REGEX, function (match) {
                //replace each letter with a star

                var stars = '';
                for (var i = 0; i < match.length; i++) {
                    stars += '*';
                }
                return stars;
            });
        }

    };
});
