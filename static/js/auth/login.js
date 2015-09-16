/* global jQuery, wisply */
var login;

/**
 * @file Encapsulates the functionality for the login page.
 * @author Cristian Sima
 */

/**
 * The constructor calls the init function.
 * @constructor
 */
function Login() {
  this.init();
}
Login.prototype = {
  /**
  * It loads the listeners and focuses the name
  */
  init : function () {
    this.loadListeners();
    this.focusName();
  },
  /**
  * It adds a listener for form submit
  */
  loadListeners: function() {
    $("#login-form").on("submit", this.FireFormSubmited);
  },
  /**
  * It focuses the name field.
  */
  focusName: function () {
    $("#login-name").focus();
  },
  /**
  * It is fired when the form has been submitted. It shows the loading image
  */
  FireFormSubmited: function() {
    wisply.showLoading('#login-submit-div', "medium");
  }
};
/**
* It is called after the page has been loaded. It creates the login object
*/
function initLogin() {
  login = new Login();
}

$(document).ready(initLogin);
