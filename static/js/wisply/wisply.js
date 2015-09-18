/* global $, bootbox, base64_decode */
/**
* It contains a reference to the Wisply.App object
* @global
* @see Wisply.App
*/
var wisply;

/**
* @file Encapsulates the functionality for all pages.
* @author Cristian Sima
*/

/**
* Encapsulates the functionality for all pages.
* @namespace Wisply
*/
var Wisply = function () {
  'use strict';

  /**
  * The constructor calls the init method
  * @class ShortcutManager
  * @memberof Wisply
  * @classdesc It manages the operations with the key shortcuts
  */
  function ShortcutManager() {
    this.defaultShortcuts = [{
      "type": "keyup",
      "key": "Alt+a",
      "callback": function () {
        wisply.goTo("/accessibility");
      }
    }, {
      "type": "keyup",
      "key": "Alt+h",
      "callback": function () {
        wisply.goTo("/");
      }
    }, {
      "type": "keyup",
      "key": "Alt+c",
      "callback": function () {
        wisply.goTo("/contact");
      }
    }];
  }
  ShortcutManager.prototype =
  /**
  * @lends Wisply.ShortcutManager
  */
  {
    /**
    * Called when the object is create. It activates the default shortcuts
    */
    init: function () {
      this.activate(this.defaultShortcuts);
    },
    /**
    * It activates the shortcuts received as parameters
    * @param  {array} shortcuts An array with the shortcuts to active. A shortcut has a event type, the shortcut combination of keys and the callback
    */
    activate: function (shortcuts) {
      var shortcut;
      for (var i = 0; i < shortcuts.length; i++) {
        shortcut = shortcuts[i];
        $(document).bind(shortcut.type, shortcut.key, shortcut.callback);
      }
    }
  };

  /**
  * It creates the object
  * @class Message
  * @memberof Wisply
  * @classdesc It uses manages the operating regarding JavaScript messages
  */
  function Message() {
  }
  /**
  * @memberof Message
  */
  Message.prototype =
  /**
  * @lends Wisply.Message
  */
  {
    /**
    * It shows a succesful message
    * @param  {string} message The content of the message to be displayed
    */
    showSuccess: function (message) {
      this.show("<div class='text-success'>Success</div>", message);
    },
    /**
    * It shows an error message
    * @param  {string} message The content of the message to be displayed
    */
    showError: function (message) {
      this.show("<div class='text-warning'>Sorry</div>", message);
    },
    /**
    * It shows a message
    * @param  {string} title   The title of the message
    * @param  {string} content The content of the message
    */
    show: function (title, content) {
      this.dialog({
        title: title,
        message: content
      });
    },
    /**
    * It represents an adapter for the bootbox alert function. It shows an error message
    * @param  {object} args The arguments for the dialog
    * @see {@link http://bootboxjs.com/|Bootbox official website}
    */
    alert: function (args) {
      bootbox.dialog(args);
    },
    /**
    * It represents an adapter for the bootbox alert function. It shows a dialog message
    * @param  {object} args The arguments for the dialog
    * @see {@link http://bootboxjs.com/|Bootbox official website}
    */
    dialog: function (args) {
      bootbox.dialog(args);
    }
  };

  /**
  * The constructor creates a message and a shortcut objects
  * @property {Message}  message                    The reference to the message object
  * @property {ShortcutManager}  shortcutManager    The reference to the shortcut manager
  * @class App
  * @memberof Wisply
  * @classdesc It represents the main object of the website. It stores references to other objects and it provides the main functions
  */
  var App = function App() {
    /**
    * @access public
    */
    this.message = new Message();
    /**
    * @access public
    */
    this.shortcutManager = new ShortcutManager();
  };
  App.prototype =
  /**
  * @lends Wisply.App
  */
  {
    /**
    * It initiates the shorcuts' manager
    */
    init: function () {
      this.shortcutManager.init();
    },
    /**
     * It preloads the loading image in the background and stores it in the browser cache
     */
    preloadLoadingImage: function() {
      var img = new Image();
        img.src = "/static/img/wisply/load.gif";
    },
    /**
    * It executes a JQuery post request, adding to it the xsrf token value
    * @param  {object} args Same arguments for as for a JQuery AJAX request
    * @see {@link http://api.jquery.com/jquery.ajax/|JQuery AJAX API}
    */
    executePostAjax: function (args) {
      if (typeof args.data === 'undefined') {
        args.data = {};
      }
      args.dataType = "text";
      args.method = args.type = "POST";
      var xsrf,
      xsrflist;
      xsrf = $.cookie("_xsrf");
      xsrflist = xsrf.split("|");
      args.data._xsrf = base64_decode(xsrflist[0]);
      $.ajax(args);
    },
    /**
    * It refreshes the page
    * @param  {number} delayTime The amount of time in ms to delay the refresh
    */
    reloadPage: function (delayTime) {
      if (typeof size === 'undefined') {
      } else {
        if(delayTime === "now") {
          delayTime = 0;
        }
      }
      setTimeout(function () {
        location.reload();
      }, delayTime);
    },
    /**
    * It transforms a HTML object in the loading icon for Wisply
    * @param  {string} idElement The id of the element
    * @param  {string} size      The size of the loading icon. It can be small (for 20px), medium (for 55px) and large (for 110px)
    */
    showLoading: function (idElement, size) {
      /**
      * It returns the dimension in pixels acording to string type
      * @param  {string} size The demension of the image. It can be small (for 20px), medium (for 55px) and large (for 110px)
      * @return {int}      The dimension in pixels
      */
      function getDimension(size) {
        var px = 0;
        switch (size) {
          case "small":
          px = 20;
          break;
          case "medium":
          px = 55;
          break;
          case "big":
          px = 110;
          break;
        }
        return px;
      }

      /**
      * It returns the HTML code for the loading element
      * @param  {number} dimension The size of the image in px
      * @return {string}           The HTML code for loading element
      */
      function getHTML(dimension) {
        return "<img src='/static/img/wisply/load.gif' style='height: " + dimension + "px; width: " + dimension + "px' />";
      }
      var HTML, dimension, element;

      if (typeof size === 'undefined') {
        size = "small";
      }
      element = $(idElement);
      dimension = getDimension(size);
      HTML = getHTML(dimension);
      element.html(HTML);
    },
    /**
    * It redirects the account to a certain page
    * @param  {string} address The address of the page
    */
    goTo: function (address) {
      document.location = address;
    },
    /**
    * It adds a connection to Wisply
    * @param {Connection} connection The Connection object
    */
    addConnection: function (connection) {
      this.connection = connection;
    }
  };
  return {
       App: App
   };
};
$(document).ready(function() {
  "use strict";
  var module = new Wisply();
  wisply = new module.App();
  wisply.init();
});
