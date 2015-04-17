define(['can', 'util', 'i18n', 'jquery'],
    function (can, util, i18n, $) {
    'use strict';

    /**
    * Control for Destroy
    * @private
    * @author dorajistyle
    * @param {string} target
    * @name partial#Destroy
    * @constructor
    */
    var Destroy = can.Control.extend({
      init: function () {
        this.isConfirmed = false;
        util.logInfo('*partial/Destroy', 'Initialized');
      },
      '.delete-{name}-confirm click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.itemId = util.getId(ev);
        this.performConfirm();
        return false;
      },
      '.delete-{name}-final click': function (el, ev) {
        util.logInfo('*partial/Destroy', 'Initialized');
        ev.preventDefault();
        ev.stopPropagation();
        this.isConfirmed = true;
        can.when(this.modal.hide()).then(this.performDestroy());
        return false;
      },
      '.cancel-confirm click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.isConfirmed = false;
        return false;
      },

      /**
      * Show confirm modal.
      * @memberof partial#Destroy
      */
      performConfirm: function () {
        this.modal = $.UIkit.modal('#'+this.options.modalDestroyId);
        this.modal.show();
      },
      /**
      * perform Destory action.
      * @memberof partial#Destroy
      */
      performDestroy: function () {
        var item = this.element.find('#'+this.options.modalDestroyId);
        var itemId = this.itemId;
        var name = this.options.name;
        var model = this.options.model;
        var view = this.options.view;
        var $form = $(item);
        if (!$form.data('submitted')) {
          $form.data('submitted', true);
          var destroyBtn = this.element.find('.delete-'+name+'-final');
          destroyBtn.attr('disabled', 'disabled');
          can.when(model.destroy(itemId)).then(function () {
            util.showSuccessMsg(i18n.t('destroy.done', can.capitalize(name)));
            destroyBtn.removeAttr('disabled');
            $form.data('submitted', false);
            view.reload();
          }, function (xhr) {
            destroyBtn.removeAttr('disabled');
            $form.data('submitted', false);
            util.showErrorMsg(i18n.t('destroy.fail', name));
          });

        }
      }
    });
    return Destroy;
});
