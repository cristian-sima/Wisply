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
      var instance = this;
      $(".deleteRepositoryButton").click(function(event) {
        event.preventDefault();
        var object,repository;
        object = $(this);
        repository = new Repository({
          id: object.data("id"),
          name: object.data("name"),
        });
        instance.confirmDelete(repository);
      });
      $(".emptyRepositoryButton").click(function(event){
        event.preventDefault();
        var object,
        repository;
        object = $(this);
        repository = new Repository({
          id: object.data("id"),
          name: object.data("name"),
        });
        instance.confirmEmpty(repository);
      });
      $(".showStatusMore").click(function(event){
        event.preventDefault();
        wisply.message.dialog({
          onEscape: true,
          title: "Repository Status",
          message: '<div style="background:white"><blockquote>The status defines the state of the repository<br />The following diagram explains how the status changes <br /></blockquote><img class="img-responsive" src="/static/img/admin/repository/status.png"  /></div>',
        });
      });
      GUI.activateActionListeners();
    },
    /**
    * It is called when the user wants to delete a repository. It asks for confirmation
    * @param  {Repository} repository The reference to the repository object
    */
    confirmDelete: function (repository) {
      var msg = this.getDeleteDialogMessage(repository);
      wisply.message.dialog(msg);
    },
    /**
    * It is called when the user wants to clear a repository. It asks for confirmation
    * @param  {Repository} repository The reference to the repository object
    */
    confirmEmpty: function (repository) {
      var msg = this.getEmptyDialogMessage(repository);
      wisply.message.dialog(msg);
    },
    /**
    * It returns the object which contain the arguments for the confirmation dialog
    * @param  {string} type The type of action: "delete", "clear"
    * @return {Object}        The arguements for dialog
    * @see http://bootboxjs.com/
    */
    getDeleteDialogMessage: function (repository) {
      var buttons,
      cancelButton,
      msg,
      mainButton;

      cancelButton = {
        label: "Cancel",
        className: "btn-success",
        callback: function () {
          this.modal('hide');
        }
      };
      mainButton = {
        label: "Delete",
        className: "btn-danger",
        callback: function () {
          wisply.repositoriesManager.delete(repository);
        }
      };
      buttons = {
        "cancel": cancelButton,
        "main": mainButton
      };
      msg = {
        title: "We need your confirmation",
        message: "The repository <b>" + repository.name + "</b> will be permanently removed. Are you sure?",
        onEscape: true,
        buttons: buttons
      };
      return msg;
    },
    /**
    * It returns the object which contain the arguments for the confirmation dialog
    * @param  {string} type The type of action: "delete", "clear"
    * @return {Object}        The arguements for dialog
    * @see http://bootboxjs.com/
    */
    getEmptyDialogMessage: function (repository) {
      var buttons,
      cancelButton,
      msg,
      mainButton;

      cancelButton = {
        label: "Cancel",
        className: "btn-success",
        callback: function () {
          this.modal('hide');
        }
      };
      mainButton = {
        label: "Clear metadata",
        className: "btn-danger",
        callback: function () {
          wisply.repositoriesManager.empty(repository);
          wisply.message.tellToWait("Removing metadata...");
        }
      };
      buttons = {
        "cancel": cancelButton,
        "main": mainButton
      };
      msg = {
        title: "We need your confirmation",
        message: "The metadata from this repository, stored on Wisply will be deleted. Also, Wisply will delete any data related to this (logs, data). <br /><br />Are you sure?",
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
        //wisply.message.showSuccess("The repository has been removed! Refreshing page...");
        window.location = "/admin/repositories";
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/admin/repositories/repository/' + repository.id + "/delete",
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    },
    /**
    * It clears the contents of a repository
    * @param  {Repository} repository The repository object
    */
    empty: function (repository) {
      var request,
      successCallback,
      errorCallback;

      /**
      * The callback called when the repository has been cleared. It shows a message and reloads the page in 2 seconds
      */
      successCallback =  function () {
        //wisply.message.showSuccess("The repository has been cleared! Refreshing page...");
        window.location = "/admin/repositories";
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/admin/repositories/repository/' + repository.id + "/empty",
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    }
  };

  /**
  * The constructor activates the listeners
  * @memberof Repositories
  * @class GUI
  * @classdesc It encapsulets the GUI functionality
  */
  var GUI = function GUI() {
  };
  /**
   * It returns the HTML span for a status
   * @param  {string} status The status of the repository
   * @return {string} The HTML code for the status
   */
  GUI.getStatusColor = function (status) {
    var html = "",
        label = "";
    switch (status) {
    case "unverified":
      label = "info";
        break;
    case "problems":
    case "verification-failed":
      label = "danger";
        break;
    case "updating":
    case "verifying":
    case "initializing":
      label = "warning";
        break;
    case "ok":
    case "verified":
      label = "success";
        break;
    default:
      console.log("The status [" + status + "] is not a valid one");
      break;
    }
    html += '<span class="label label-' + label + '">' + status + '</span>';
    return html;
  };
  /**
   * It activates all the listeners for the actions
   */
  GUI.activateActionListeners = function() {
    $(".repositories-init-harvest").click(function(event){
      event.preventDefault();
      var object,
      id,
      xsrf;
      object = $(this);
      id = object.data("id");
      xsrf = wisply.getXSRF();
      $('<form action="/admin/harvest/init/' + id + '" method="POST">' +
          '<input type="hidden" name="_xsrf" value="' + xsrf + '">' +
          '</form>').submit();
    });
    $(function () {
      $('[data-toggle="tooltip"]').tooltip();
    });
  };
  return {
    Repository: Repository,
    Manager: Manager,
    GUI: GUI
  };
};
$(document).ready(function() {
  "use strict";
  var module = new Repositories();
  wisply.repositoriesModule = module;
  wisply.repositoriesManager = new module.Manager();
  wisply.repositoriesManager.init();
});
