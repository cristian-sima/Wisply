/* global jQuery, bootbox, base64_decode */
var wisply;

(function ($) {
  'use strict';

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

  function Shortcut() {
    this.init();
  }

  Shortcut.prototype = {
    init: function () {
      this.activateShortcuts(defaultShortcuts);
    },
    activateShortcuts: function (shortcuts) {
      var shortcut;
      for (var i = 0; i < shortcuts.length; i++) {
        shortcut = shortcuts[i];
        $(document).bind(shortcut.type, shortcut.key, shortcut.callback);
      }
    }
  };

  function Message() {

  }

  Message.prototype = {
    showSuccess: function (message) {
      this.show("<div class='text-success'>Success</div>", message);
    },
    showError: function (message) {
      this.show("<div class='text-warning'>Sorry</div>", message);
    },
    show: function (title, content) {
      bootbox.dialog({
        title: title,
        message: content
      });
    },
    alert: function (args) {
      bootbox.dialog(args);
    },
    dialog: function (args) {
      bootbox.dialog(args);
    }
  };

  function Wisply() {
    this.message = new Message();
    this.shortcut = new Shortcut();
  }
  Wisply.prototype = {
    executePostAjax: function (args) {
      if (typeof args.data === 'undefined') {
        args.data = {};
      }
      args.dataType = "text";
      args.method = "POST";
      args.type = "POST";
      var xsrf,
      xsrflist;
      xsrf = $.cookie("_xsrf");
      xsrflist = xsrf.split("|");
      args.data._xsrf = base64_decode(xsrflist[0]);
      $.ajax(args);
    },
    reloadPage: function () {
      setTimeout(function () {
        location.reload();
      }, 2000);
    },
    showLoading: function (idElement, size) {
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
    goTo: function (address) {
      document.location = address;
    },
    addConnection: function (connection) {
      this.connection = connection;
    }
  };

  function initPage() {
    wisply = new Wisply();
    $('[data-toggle=offcanvas], #close-sidebar-admin').click(function() {
      $('.row-offcanvas').toggleClass('active');
      $('html, body').animate({ scrollTop: 0 }, 'fast');
    });
  }
  $(document).ready(initPage);
}(jQuery, wisply, bootbox));
