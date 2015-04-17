define(['can',
        'util', 'validation', 'i18n', 'jquery'],
    function (can, util, validation, i18n, $) {
    'use strict';

    var Sharing =  can.Control.extend({
      init: function () {
        util.logInfo('*partial/Share', 'Initialized');
      },
      //        '.fb-like click': function (el, ev) {
      //            var url = util.getCurrentURL();
      //            var host = util.getCurrentHost();
      //            var title = util.getData(ev,'title');
      //            var description = '';
      //            var $img = '';
      //            var picture = '';
      //            var $pageDescription = $('#pageDescription');
      //            if(url == undefined) { url = '';}
      //            if(host == undefined) { host = '';}
      //            if($pageDescription != null) {
      //                description = $pageDescription.text();
      //                $img = $($pageDescription.find('img'));}
      ////            var description = util.getData(ev,'description');
      //            var tags = util.getData(ev,'tags');
      //            util.logDebug("url", url);
      ////            var picture = '';
      ////            if(picture.indexOf("://") < 0) {
      ////                picture = host+picture;
      ////            }
      ////            var pictureYoutube = http://img.youtube.com/vi/bWGqmOp4wmk/0.jpg
      ////            var pictureVimeo = http://vimeo.com/api/v2/video/101952159.json
      ////            picture = decodeURIComponent(picture);
      //            var meta = '<meta property="og:title" content="'+title+'" />'+
      //                       '<meta property="og:description" content="'+description+'" />';
      //            if($img.length > 0) {
      //                picture = $img.prop('src');
      //                util.logDebug('img', picture);
      //                meta+='<meta property="og:image" content="'+picture+'" />';
      //            }
      //
      ////              FB.ui({
      ////             method: 'share_open_graph',
      ////             action_type: 'og.likings',
      ////             action_properties: JSON.stringify({
      ////              object: url
      ////             })
      ////            }, function(response){
      ////                    if(response != undefined) {
      ////                   $('#fbCount').html(util.fetchFBLikingCount(url));
      ////                   util.showInfoMsg(i18n.t('article.view.item.facebook.share.done'));
      ////                  }
      ////                  util.logDebug('share fb',response);
      ////
      ////              });
      //
      //
      //            FB.api(
      //                "/me/og.likings",
      //                "POST",
      //                {
      //                    "object": {
      //                        "og:title": title,
      //                        "og:url": url.replace('/','\/'),
      //                        "og:image": picture.replace('/','\/'),
      //                        "object":url.replace('/','\/')
      //                    }
      //                },
      //                function (response) {
      //                   if(response != undefined) {
      //                       util.fetchFBLikingCount(url).always(function(res){
      //                            $('#fbCount').text(res);
      //                        });
      //                   util.showInfoMsg(i18n.t('article.view.item.facebook.share.done'));
      //                  }
      //                  util.logDebug('share fb',response);
      //
      //
      //                  if (response && !response.error) {
      //                    /* handle the result */
      //                  }
      //                }
      //            );
      //
      //        },
      '.share-to-fb click': function (el, ev) {
        var url = util.getCurrentURL();
        var host = util.getCurrentHost();
        var title = util.getData(ev,'title');
        var description = '';
        var $img = '';
        var picture = '';
        var $pageDescription = $('#pageDescription');
        if(url == undefined) { url = '';}
        if(host == undefined) { host = '';}
        if($pageDescription != null) {
          description = $pageDescription.text();
          $img = $($pageDescription.find('img'));}
          var tags = util.getData(ev,'tags');

          //            var picture = '';
          //            if(picture.indexOf("://") < 0) {
          //                picture = host+picture;
          //            }
          //            var pictureYoutube = 'http://img.youtube.com/vi/.jpg'
          //            var pictureVimeo = 'http://vimeo.com/api/v2/video/.json'
          //            picture = decodeURIComponent(picture);
          var meta = '<meta property="og:title" content="'+title+'" />'+
          '<meta property="og:description" content="'+description+'" />';
          if($img.length > 0) {
            picture = $img.prop('src');
            util.logDebug('img', picture);
            meta+='<meta property="og:image" content="'+picture+'" />';
          }

          var mediaArray = [];
          if(picture.length >0) {
            mediaArray = [{"type":"image","src": picture,"href": url}];
          }
          FB.ui({
            method: "stream.publish",
            display: "iframe",
            privacy: {"value": "EVERYONE"},
            //                user_message_prompt: "Publish the article!",
            //                message: "I am so smart!  S M R T!",
            attachment: {
              name: title,
              caption: host,
              href: url,
              media: mediaArray

              //                   properties:{
              //                     "1)":{"text":"goyangi","href": host}
              //                   }
            }
            //                action_links: [{ text: tag.name, href: tag.link}]
          }, function(response){
            if(response != undefined) {
              util.showInfoMsg(i18n.t('article.view.item.facebook.share.done'));
            }
            util.logDebug('share fb',response);
          });

          util.logDebug('share fb title',title);
          util.logDebug('share fb description', description);
          //            window.open("https://www.facebook.com/sharer/sharer.php?u="+url);

          //            FB.ui({
          //              method: 'feed',
          //              link: url,
          ////              link: 'http://goyangi.dorajistyle.github.io/',
          //              privacy: {"value": "EVERYONE"},
          //              name: title,
          ////              caption: host,
          //              description: description,
          //              picture: picture
          ////              actions:
          //            }, function(response){
          //                if(response != undefined) {
          //                   util.showInfoMsg(i18n.t('article.view.item.facebook.share.done'));
          //                }
          //                util.logDebug('share fb',response);
          //            });
          //            FB.ui({ method: 'apprequests',
          //                message: 'Request to use my app'
          //            }, function(response){
          //                util.showInfoMsg('apprequests success');
          //            });
        }
      });
      return Sharing;

});
