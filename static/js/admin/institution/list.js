/* global $, wisply */

/**
* @file Encapsulates the functionality for managing the institutions
* @author Cristian Sima
*/

/**
* @namespace Institutions
*/
var Institutions = function () {
  'use strict';

  /**
  * The constructor does nothing important
  * @class Institution
  * @memberof Institutions
  * @classdesc It represents a institution
  * @param {object} info It contains the information regarding the institution (id, name and web address)
  */
  var Institution = function Institution(info) {
    this.id = info.id;
    this.name = info.name;
    this.url = info.url;
  };

  /**
  * The constructor activates the listeners
  * @memberof Institutions
  * @class Manager
  * @classdesc It encapsulets the functionality for the institutions
  */
  var Manager = function Manager() {
  };
  Manager.prototype =
  /** @lends Institutions.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.activateListeners();
    },
    /**
    * It activates the listener for deleting a institution
    * @fires InstitutionsManager#confirmDelete
    */
    activateListeners: function () {
      $(".deleteInstitutionButton").click(confirmDelete);
      GUI.activateActionListeners();
    },
    /**
    * It is called when the user wants to delete a institution. It asks for confirmation
    * @param  {Institution} institution The reference to the institution object
    */
    confirmDelete: function (institution) {
      var msg = this.getDialogMessage(institution);
      wisply.message.dialog(msg);
    },
    /**
    * It returns the object which contain the arguments for the confirmation dialog
    * @param  {Institution} institution The institution object
    * @return {Object}        The arguements for dialog
    * @see http://bootboxjs.com/
    */
    getDialogMessage: function (institution) {
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
          wisply.institutionsManager.delete(institution);
        }
      };
      buttons = {
        "cancel": cancelButton,
        "main": mainButton
      };
      msg = {
        title: "Please confirm!",
        message: "The institution <b>" + institution.name + "</b> will be permanently removed. Are you sure?",
        onEscape: true,
        buttons: buttons
      };
      return msg;
    },
    /**
    * It delets a institution
    * @param  {Institution} institution The institution object
    */
    delete: function (institution) {
      var request,
      successCallback,
      errorCallback;

      /**
      * The callback called when the institution has been deleted. It shows a message and reloads the page in 2 seconds
      */
      successCallback =  function () {
        wisply.message.showSuccess("The institution has been removed! Refreshing page...");
        wisply.reloadPage(2000);
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/admin/institutions/delete/' + institution.id,
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    }
  };

  /**
  * It is called when the user clicks the delete button. It creates the institution button and asks for confirmation
  * @fires InstitutionsManager#confirmDelete()
  * @param  {event} e The event generated
  */
  function confirmDelete(e) {
    e.preventDefault();
    var instance,
    name,
    id,
    institution;
    instance = $(this);
    id = instance.data("id");
    name = instance.data("name");
    institution = new Institution({
      id: id,
      name: name
    });
    wisply.institutionsManager.confirmDelete(institution);
  }

  function initInstitution(e) {
      e.preventDefault();
      var instance,
      id,
      xsrf;
      instance = $(this);
      id = instance.data("id");
      xsrf = wisply.getXSRF();

        $('<form action="/admin/harvest/init/' + id + '" method="POST">' +
          '<input type="hidden" name="_xsrf" value="' + xsrf + '">' +
          '</form>').submit();

  }

  /**
  * The constructor activates the listeners
  * @memberof Institutions
  * @class GUI
  * @classdesc It encapsulets the GUI functionality
  */
  var GUI = function GUI() {
  };
  /**
   * It activates all the listeners for the actions
   */
  GUI.activateActionListeners = function() {
    $(".institutions-init-harvest").click(initInstitution);
    $(function () {
      $('[data-toggle="tooltip"]').tooltip();
    });
  };
  return {
    Institution: Institution,
    Manager: Manager,
    GUI: GUI
  };
};
$(document).ready(function() {
  "use strict";
  var module = new Institutions();
  wisply.institutionsModule = module;
  wisply.institutionsManager = new module.Manager();
  wisply.institutionsManager.init();
});
