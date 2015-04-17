define(['i18n'],
function (i18n) {
  'use strict';
  var HTTP404    = "404",       // do not load any image if none is associated with the email hash, instead return an HTTP 404 (File Not Found) response
  MYSTERYMAN = "mm",        // (mystery-man) a simple, cartoon-style silhouetted outline of a person (does not vary by email hash)
  IDENTICON  = "identicon", // a geometric pattern based on an email hash
  MONSTERID  = "monsterid", // a generated 'monster' with different colors, faces, etc
  WAVATAR    = "wavatar",   // generated faces with differing features and backgrounds
  RETRO      = "retro",     // awesome generated, 8-bit arcade-style pixelated faces
  BLANK      = "blank",   // a transparent PNG image (border added to HTML below for demonstration purposes)
  PREFIX     = "www",
  DEFAULT_SIZE = 100;
  return {
    getGravatarURL: function (hash,size,defaultIcon, secure) {
      if(size == undefined) {
        size = DEFAULT_SIZE;
      }
      if(defaultIcon == undefined) {
        defaultIcon = MYSTERYMAN;
      }
      if(secure) {
        PREFIX = "secure";
      }
      return "http://"+PREFIX+".gravatar.com/avatar/"+hash+"?s="+size+"&d="+defaultIcon;
    }
  }
});
