/* global jQuery, wisply */
var accounts;

(function ($) {
    'use strict';

    /**
     * It represents an account
     * @param {number} id   The id of the account
     * @param {string} name The name of the account
     */
    function Account(id, name) {
        this.id = id;
        this.name = name;
    }

    /**
     * It encapsulets the functions for accounts.
     * The constructor activates the listeners
     */
    function Accounts() {
        this.activateListeners();
    }
    Accounts.prototype = {
        /**
         * It activates the listener for all delete buttons
         */
        activateListeners: function () {
            $(".deleteAccountButton").click(confirmDelete);
        },
        /**
         * It requests the user to confirm
         * @param  {Account} account  The Account object to be deleted
         */
        confirmDelete: function (account) {
            var msg = this.getDialogMessage(account);
            wisply.message.dialog(msg);
        },
        /**
         * It returns the arguements of the confimation message
         * @param  {Account} account The Account object to be deleted
         * @return {object}         The arguements of the confimation message
         */
        getDialogMessage: function (account) {
            var buttons,
                cancelButton,
                msg,
                mainButton;

            cancelButton = {
                label: "Cancel",
                className: "btn-default",
                callback: function () {
                    this.modal('hide');
                }
            };
            mainButton = {
                label: "Delete",
                className: "btn-primary",
                callback: function () {
                    accounts.delete(account);
                }
            };
            buttons = {
                "cancel": cancelButton,
                "main": mainButton
            };
            msg = {
                title: "Please confirm!",
                message: "The account <b>" + account.name + "</b> will be permanently removed. Are you sure?",
                onEscape: true,
                buttons: buttons
            };
            return msg;
        },
        /**
         * It delets a user
         * @param  {Account} account The account to be deleted
         */
        delete: function (account) {
          var request,
          successCallback,
          errorCallback;

          /**
           * It is called when the deletion has been done. It reloads the page in 2 seconds
           * @return {[type]} [description]
           */
          successCallback =  function () {
              wisply.message.showSuccess("The account has been removed! Refreshing page...");
              wisply.reloadPage();
          };

          /**
           * It is called when there has been problems
           */
          errorCallback = function () {
              wisply.message.showError("There was a problem with your request!");
          };

          request = {
            "url": '/admin/accounts/delete/' + account.id,
            "success": successCallback,
            "error": errorCallback
          };
          wisply.executePostAjax(request);
        }
    };

    /**
     * It is called when the user clicks the delete button. It requests the user to confirm
     * @param  {Event} event The event generated
     */
    function confirmDelete(event) {
        e.preventDefault();
        var instance,
            name,
            id,
            account;
        instance = $(this);
        id = instance.data("id");
        name = instance.data("name");
        account = new Account(id, name);
        accounts.confirmDelete(account);
    }

    /**
     * It is called when the page has loaded. It creates tha Accounts object
     */
    function initAccounts() {
        accounts = new Accounts();
    }
    $(document).ready(initAccounts);
}(jQuery, wisply, accounts));
