/* global jQuery, wisply, bootbox*/
var register;

(function ($) {
  'use strict';
  /**
  * It encapsulates the functionality for the register page
  */
  function Register() {
    this.init();
  }
  Register.prototype = {
    /**
    * It activates the listeners and focuses the name
    */
    init : function () {
      this.loadListeners();
      this.checkConfirmPassword();
      this.focusName();
    },
    /**
     * It activates the listeners for form submitted and password focused
     */
    loadListeners: function() {
      this.passwordCompletedListener();
      this.formSubmittedListener();
    },
    /**
     * It is called when the user clicks on the back button of the browser (or in the case Wisply detects problems with the form and the user goes back to the form). In case the value of the password is not empty, it shows the confirmation password field
     */
    checkConfirmPassword : function () {
      if($("#register-password-confirm").val() !== "") {
        $("#div-confirm-password").show();
      }
    },
    /**
    * It focuses the name field
    */
    focusName: function () {
      $("#register-name").focus();
    },
    /**
    * It is called when the form has been submitted. It shows the loading button
    */
    formSubmittedListener : function() {
      $("#register-form").on("submit", this.FireFormSubmited);
    },
    /**
     * It is called when the register form has been submitted. It checks if the confirmation password is the same as the password. If so, it submits the form, else it shows a message
     * @param  {Event} event The event which is generated
     */
    FireFormSubmited: function(event) {
      event.preventDefault();
      var password = $('#register-password').val(),
      confirmationPassword = $("#register-password-confirm").val();
      if (password === confirmationPassword) {
        register.showLoading();
        this.submit();
      } else {
        register.showPasswordsDoNotMatch();
      }
    },
    /**
     * It shows the confirmation password field
     */
    passwordCompletedListener : function () {
      $("#register-password").focus(function() {
        $("#div-confirm-password").show();
      });
    },
    /**
     * It shows the loading image
     */
    showLoading: function() {
      wisply.showLoading('#register-submit-div', "medium");
    },
    /**
     * It tells the account that the passwords do not match
     */
    showPasswordsDoNotMatch: function () {
      var args = {
        title: "I got a problem",
        message: "The confirmation password is equal to the password. Correct this.",
        onEscape: true,
        buttons: {
          cancel: {
            label: "Ok",
            className: "btn-default",
            callback: function() {
              this.modal('hide');
            }
          }}
        };
        wisply.message.alert(args);
      }
    };
    /**
    * It is called when the page has been loaded. It creates the register button
    */
    function initRegister() {
      register = new Register();
    }
    $(document).ready(initRegister);
  }(jQuery, wisply, bootbox, register));
