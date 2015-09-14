/* global jQuery, wisply */
var accounts;

(function ($) {
    'use strict';

    function Account(id, name) {
        this.id = id;
        this.name = name;
    }

    function Accounts() {
        this.activateListeners();
    }
    Accounts.prototype = {
        activateListeners: function () {
            $(".deleteAccountButton").click(confirmDelete);
        },
        confirmDelete: function (account) {
            var msg = this.getDialogMessage(account);
            wisply.message.dialog(msg);
        },
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
        delete: function (account) {
          var request,
          successCallback,
          errorCallback;

          successCallback =  function () {
              wisply.message.showSuccess("The account has been removed! Refreshing page...");
              wisply.reloadPage();
          };

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

    function confirmDelete(e) {
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

    function initAccounts() {
        accounts = new Accounts();
    }
    $(document).ready(initAccounts);
}(jQuery, wisply, accounts));
