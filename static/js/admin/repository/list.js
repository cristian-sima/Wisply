/* global $, wisply */

/**
* @file Encapsulates the functionality for managing the repositories
* @author Cristian Sima
*/


/**
* @namespace Repositories
*/
var Repositories = function () {
  'use strict';

  /**
  * The constructor does nothing important
  * @class Repository
  * @memberof Repositories
  * @classdesc It represents a repository
  * @param {object} info It contains the information regarding the repository (id, name and url, status)
  */
  var Repository = function Repository(info) {
    this.id = info.id;
    this.name = info.name;
    this.url = info.url;
    this.status = info.status;
  };

  /**
  * The constructor activates the listeners
  * @memberof Repositories
  * @class Manager
  * @classdesc It encapsulets the functionality for the repositories
  */
  var Manager = function Manager() {
  };
  Manager.prototype =
  /** @lends Repositories.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.activateListeners();
    },
    /**
    * It activates the listener for deleting a repository
    * @fires RepositoriesManager#confirmDelete
    */
    activateListeners: function () {
      $(".deleteRepositoryButton").click(confirmDelete);
      $(".repositories-init-harvest").click(initRepository);
      $(function () {
        $('[data-toggle="tooltip"]').tooltip();
      });
    },
    /**
    * It is called when the user wants to delete a repository. It asks for confirmation
    * @param  {Repository} repository The reference to the repository object
    */
    confirmDelete: function (repository) {
      var msg = this.getDialogMessage(repository);
      wisply.message.dialog(msg);
    },
    /**
    * It returns the object which contain the arguments for the confirmation dialog
    * @param  {Repository} repository The repository object
    * @return {Object}        The arguements for dialog
    * @see http://bootboxjs.com/
    */
    getDialogMessage: function (repository) {
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
          wisply.repositoriesManager.delete(repository);
        }
      };
      buttons = {
        "cancel": cancelButton,
        "main": mainButton
      };
      msg = {
        title: "Please confirm!",
        message: "The repository <b>" + repository.name + "</b> will be permanently removed. Are you sure?",
        onEscape: true,
        buttons: buttons
      };
      return msg;
    },
    /**
    * It delets a repository
    * @param  {Repository} repository The repository object
    */
    delete: function (repository) {
      var request,
      successCallback,
      errorCallback;

      /**
      * The callback called when the repository has been deleted. It shows a message and reloads the page in 2 seconds
      */
      successCallback =  function () {
        wisply.message.showSuccess("The repository has been removed! Refreshing page...");
        wisply.reloadPage(2000);
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/admin/repositories/delete/' + repository.id,
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    }
  };

  /**
  * It is called when the user clicks the delete button. It creates the repository button and asks for confirmation
  * @fires RepositoriesManager#confirmDelete()
  * @param  {event} e The event generated
  */
  function confirmDelete(e) {
    e.preventDefault();
    var instance,
    name,
    id,
    repository;
    instance = $(this);
    id = instance.data("id");
    name = instance.data("name");
    repository = new Repository({
      id: id,
      name: name
    });
    wisply.repositoriesManager.confirmDelete(repository);
  }

  function initRepository(e) {
      e.preventDefault();
      var instance,
      id,
      xsrf,
      repository;
      instance = $(this);
      id = instance.data("id");
      xsrf = wisply.getXSRF();

        $('<form action="/admin/harvest/init/' + id + '" method="POST">' +
          '<input type="hidden" name="_xsrf" value="' + xsrf + '">' +
          '</form>').submit();

  }

  return {
    Repository: Repository,
    Manager: Manager
  };
};
$(document).ready(function() {
  "use strict";
  var module = new Repositories();
  wisply.repositoriesManager = new module.Manager();
  wisply.repositoriesManager.init();
});
