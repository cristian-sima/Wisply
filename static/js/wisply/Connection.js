/* global $, wisply */

/**
* @file Encapsulates the functionality for managing the connection of an account (both user and administrator)
* @author Cristian Sima
*/

/**
* @namespace Connection
*/
var Connection = function () {
  'use strict';


  /**
  * The constructor calls the init function
  * @class Connection
  * @memberof Connection
  * @classdesc The object manages the connection for an account.
  */
  var Connection = function Connection() {
  };
  Connection.prototype =
  /** @lends Connection.Connection */
  {
    /**
    * It loads the listeners
    */
    init: function () {
      this.loadListeners();
      this.loadShortcuts();
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
      wisply.shortcutManager.activate(shortcuts);
    },
    /**
    * It activates the listener for form submited
    * @fires Connection#FireLogoutUser
    */
    loadListeners: function () {
      $("#menu-logout-button").click(this.FireLogoutUser);
    },
    /**
    * It is called when the account wants to log out. It logs the account out
    * @event Connection#FireLogoutUser
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
  return {
    Connection: Connection
  };
};
$(document).ready(function() {
  "use strict";
  var module = new Connection();
  wisply.connection = new module.Connection();
  wisply.connection.init();
});
