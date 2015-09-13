/* global jQuery, wisply */
var sources;

(function ($) {
    'use strict';

    function Source(id, name) {
        this.id = id;
        this.name = name;
    }

    function Sources() {
        this.activateListeners();
    }
    Sources.prototype = {
        activateListeners: function () {
            $(".deleteSourceButton").click(confirmDelete);
        },
        confirmDelete: function (source) {
            var msg = this.getDialogMessage(source);
            wisply.message.dialog(msg);
        },
        getDialogMessage: function (source) {
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
                    sources.delete(source);
                }
            };
            buttons = {
                "cancel": cancelButton,
                "main": mainButton
            };
            msg = {
                title: "Please confirm!",
                message: "The source <b>" + source.name + "</b> will be permanently removed. Are you sure?",
                onEscape: true,
                buttons: buttons
            };
            return msg;
        },
        delete: function (source) {
          var request,
          successCallback,
          errorCallback;

          successCallback =  function () {
              wisply.message.showSuccess("The source has been removed! Refreshing page...");
              wisply.reloadPage();
          };

          errorCallback = function () {
              wisply.message.showError("There was a problem with your request!");
          }

          request = {
            "url": '/admin/sources/delete/' + source.id,
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
            source;
        instance = $(this);
        id = instance.data("id");
        name = instance.data("name");
        source = new Source(id, name);
        sources.confirmDelete(source);
    }

    function initSources() {
        sources = new Sources();
    }
    $(document).ready(initSources);
}(jQuery, wisply, sources));
