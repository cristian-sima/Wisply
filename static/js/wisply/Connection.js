/* global jQuery, wisply */
(function ($) {
  'use strict';

  function Connection() {
    this.init();
  }
  Connection.prototype = {
    init: function () {
      this.loadListeners();
      this.loadShortcuts();
    },
    loadListeners: function () {
      $("#menu-logout-button").click(this.FireLogoutUser);
    },
    loadShortcuts: function () {
      var shortcuts = [{
        "type": "keyup",
        "key": "Alt+l",
        "callback": function () {
          wisply.connection.logout();
        }
      }];
      wisply.shortcut.activateShortcuts(shortcuts);
    },
    FireLogoutUser: function (event) {
      wisply.showLoading("#menu-top-left", "small");
      event.preventDefault();
      wisply.connection.logout();
    },
    logout: function () {
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

      request = this.getRequestData(successCallback, errorCallback);
      wisply.executePostAjax(request);
    },
    getRequestData: function (successCallback, errorCallback) {
      var request = {
        "url": '/auth/logout',
        "dataType": "text",
        'method': "POST",
        "type": "POST",
        "success": successCallback,
        "error": errorCallback
      };
      return request;
    }
  };

  function initConnection() {
    var connection;
    connection = new Connection();
    wisply.addConnection(connection);
  }
  $(document).ready(initConnection);
}(jQuery, wisply));
