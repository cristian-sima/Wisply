/* global jQuery, wisply */
var users;

(function ($) {
    'use strict';

    function User(id, name) {
        this.id = id;
        this.name = name;
    }

    function Users() {
        this.activateListeners();
    }
    Users.prototype = {
        activateListeners: function () {
            $(".deleteUserButton").click(confirmDelete);
        },
        confirmDelete: function (user) {
            var msg = this.getDialogMessage(user);
            wisply.message.dialog(msg);
        },
        getDialogMessage: function (user) {
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
                    users.delete(user);
                }
            };
            buttons = {
                "cancel": cancelButton,
                "main": mainButton
            };
            msg = {
                title: "Please confirm!",
                message: "The user <b>" + user.name + "</b> will be permanently removed. Are you sure?",
                onEscape: true,
                buttons: buttons
            };
            return msg;
        },
        delete: function (user) {
          var request,
          successCallback,
          errorCallback;

          successCallback =  function () {
              wisply.message.showSuccess("The user has been removed! Refreshing page...");
              wisply.reloadPage();
          };

          errorCallback = function () {
              wisply.message.showError("There was a problem with your request!");
          }

          request = {
            "url": '/admin/users/delete/' + user.id,
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
            user;
        instance = $(this);
        id = instance.data("id");
        name = instance.data("name");
        user = new User(id, name);
        users.confirmDelete(user);
    }

    function initUsers() {
        users = new Users();
    }
    $(document).ready(initUsers);
}(jQuery, wisply, users));
