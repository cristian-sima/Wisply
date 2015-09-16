/* global jQuery, bootbox, base64_decode */
var wisply;


/**
* The default shortcuts for all the pages
*/
var defaultShortcuts = [{
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

/**
* It manages the operations with the key shortcuts
* @namespace Shortcut
*/
function Shortcut() {
  this.init();
}

Shortcut.prototype = {
  /**
  * Called when the object is create. It activates the default shortcuts
  */
  init: function () {
    this.activateShortcuts(defaultShortcuts);
  },
  /**
  * It activates the shortcuts received as parameters
  * @param  {array} shortcuts An array with the shortcuts to active. A shortcut has a event type, the shortcut combination of keys and the callback
  */
  activateShortcuts: function (shortcuts) {
    var shortcut;
    for (var i = 0; i < shortcuts.length; i++) {
      shortcut = shortcuts[i];
      $(document).bind(shortcut.type, shortcut.key, shortcut.callback);
    }
  }
};

/**
* It uses manages the operating regarding JavaScript messages
* @namespace Message
*/
function Message() {
}

Message.prototype = {
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
* It represents the main object of the website. It stores references to other objects and it provides the main functions
* The constructor creates a message and a shortcut objects
* @namespace Wisply
*/
function Wisply() {
  this.message = new Message();
  this.shortcut = new Shortcut();
}
Wisply.prototype = {
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
    }, delaytime);
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

/**
* It is called when the page has been loaded. It creates the Wisply object
*/
function initPage() {
  wisply = new Wisply();
}
$(document).ready(initPage);
