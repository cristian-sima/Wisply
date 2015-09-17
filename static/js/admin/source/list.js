/* global $, wisply */

/**
* @file Encapsulates the functionality for managing the sources
* @author Cristian Sima
*/


/**
* @namespace Sources
*/
var Sources = function () {
  'use strict';

  /**
  * The constructor does nothing important
  * @class Source
  * @memberof Sources
  * @classdesc It represents a source
  * @param {number} id   The id of the source
  * @param {string} name The name of the source
  */
  var Source = function Source(id, name) {
    this.id = id;
    this.name = name;
  };

  /**
  * The constructor activates the listeners
  * @memberof Sources
  * @class Manager
  * @classdesc It encapsulets the functionality for the sources
  */
  var Manager = function Manager() {
  };
  Manager.prototype =
  /** @lends Sources.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.activateListeners();
    },
    /**
    * It activates the listener for deleting a source
    * @fires SourcesManager#confirmDelete
    */
    activateListeners: function () {
      $(".deleteSourceButton").click(confirmDelete);
    },
    /**
    * It is called when the user wants to delete a source. It asks for confirmation
    * @param  {Source} source The reference to the source object
    */
    confirmDelete: function (source) {
      var msg = this.getDialogMessage(source);
      wisply.message.dialog(msg);
    },
    /**
    * It returns the object which contain the arguments for the confirmation dialog
    * @param  {Source} source The source object
    * @return {Object}        The arguements for dialog
    * @see http://bootboxjs.com/
    */
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
          wisply.sourcesManager.delete(source);
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
    /**
    * It delets a source
    * @param  {Source} source The source object
    */
    delete: function (source) {
      var request,
      successCallback,
      errorCallback;

      /**
      * The callback called when the source has been deleted. It shows a message and reloads the page in 2 seconds
      */
      successCallback =  function () {
        wisply.message.showSuccess("The source has been removed! Refreshing page...");
        wisply.reloadPage(2000);
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/admin/sources/delete/' + source.id,
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    }
  };

  /**
  * It is called when the user clicks the delete button. It creates the source button and asks for confirmation
  * @fires SourcesManager#confirmDelete()
  * @param  {event} e The event generated
  */
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
    wisply.sourcesManager.confirmDelete(source);
  }
  return {
    Source: Source,
    Manager: Manager
  };
};
$(document).ready(function() {
  "use strict";
  var module = new Sources();
  wisply.sourcesManager = new module.Manager();
  wisply.sourcesManager.init();
});
