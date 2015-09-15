/* global jQuery, wisply */
(function ($) {
  'use strict';

  /**
  * The object manages the connection for an account.
  * The constructor calls the init function
  */
  function Connection() {
    this.init();
  }
  Connection.prototype = {
    /**
    * It loads the listeners
    */
    init: function () {
      this.loadListeners();
      this.loadShortcuts();
    },
    /**
    * It activates the listener for form submited
    */
    loadListeners: function () {
      $("#menu-logout-button").click(this.FireLogoutUser);
    },
    /**
    * It loads the default shortcuts for connection
    */
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
    /**
    * It is called when the account wants to log out. It logs the account out
    * @param {Event} event The event which has been cautch
    */
    FireLogoutUser: function (event) {
      wisply.showLoading("#menu-top-left", "small");
      event.preventDefault();
      wisply.connection.logout();
    },
    /**
    * It creates and sends the AJAX request to log out.
    */
    logout: function () {
      var request,
      successCallback,
      errorCallback;

      /**
      * It returns the details of the request for log out
      * @param  {function} successCallback The callback called when the user has been logged out
      * @param  {function} errorCallback   The callback called when the log out was not possbile
      * @return {object}                 The object which contains the arguments for the request
      */
      function getRequestData(successCallback, errorCallback) {
        var request = {
          "url": '/auth/logout',
          "success": successCallback,
          "error": errorCallback
        };
        return request;
      }

      /**
      * The callback called when the user has been logged out
      * @return {[type]} [description]
      */
      successCallback = function () {
        wisply.message.showSuccess("You have been disconnected! Refreshing page...");
        wisply.reloadPage(2000);
      };

      /**
      * The callback called when the log out was not possbile
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = getRequestData(successCallback, errorCallback);
      wisply.executePostAjax(request);
    }
  };
  /**
   * It is called when the page has been loaded. It creates the connection object
   */
  function initConnection() {
    wisply.connection = new Connection();
  }
  $(document).ready(initConnection);
}(jQuery, wisply));
