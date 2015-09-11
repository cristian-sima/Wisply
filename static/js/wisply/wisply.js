/* global jQuery, bootbox, base64_decode */
var wisply;

(function ($) {
  'use strict';

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
    alert : function(args) {
      bootbox.dialog(args);
    }
  };

  function Wisply() {
    this.message = new Message();
    this.loadListeners();
  }
  Wisply.prototype = {
    loadListeners: function () {
      $("#menu-logout-button").click(this.FireLogoutUser);
    },
    FireLogoutUser: function (event) {
      wisply.showLoading("#menu-top-left", "small");
      event.preventDefault();
      wisply.tryLogout();
    },
    tryLogout: function () {
      var request,
      successCallback,
      errorCallback;

      successCallback = function () {
        wisply.message.showSuccess("You have been disconnected! Refreshing page...");
        wisply.reloadPage();
      };

      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/auth/logout',
        "dataType": "text",
        'method': "POST",
        "type": "POST",
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    },
    executePostAjax: function (args) {
      if (typeof args.data === 'undefined') {
        args.data = {};
      }
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
    }
  };

  function initPage() {
    wisply = new Wisply();
  }
  $(document).ready(initPage);
}(jQuery, wisply, bootbox));
