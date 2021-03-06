/* global $, wisply */

/**
* @file Encapsulates the functionality for managing the institutions
* @author Cristian Sima
*/

/**
* @namespace InstitutionAdvanceOptionsModule
*/
var InstitutionAdvanceOptionsModule = function () {
  'use strict';

  /**
  * The constructor does nothing important
  * @class Institution
  * @memberof InstitutionAdvanceOptionsModule
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
  * @memberof InstitutionAdvanceOptionsModule
  * @class Manager
  * @classdesc It encapsulets the functionality for the institutions
  */
  var Manager = function Manager() {
  };
  Manager.prototype =
  /** @lends InstitutionAdvanceOptionsModule.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.activateListeners();
    },
    /**
    * It activates the listener for deleting a institution
    * @fires InstitutionAdvanceOptionsModule#confirmDelete
    */
    activateListeners: function () {
      var instance = this;
      $(".deleteInstitutionButton").click(function(event){
        event.preventDefault();
        var object,
        institution;
        object = $(this);
        institution = new Institution({
          id: object.data("id"),
          name: object.data("name"),
        });
        instance.confirmDelete(institution);
      });
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
      mainButton,
      instance = this;

      cancelButton = {
        label: "No, thanks",
        className: "btn-success",
        callback: function () {
          this.modal('hide');
        }
      };
      mainButton = {
        label: "Delete",
        className: "btn-danger",
        callback: function () {
          instance.delete(institution);
        }
      };
      buttons = {
        "cancel": cancelButton,
        "main": mainButton
      };
      msg = {
        title: "We need your confirmation!",
        message: "<h2 class='text-warning'>Warning:</h2> <div>These will be permanently removed:</div> <ul>                  <li> The details of institution</li>                  <li> All the repositories                  <ul>                    <li>All the records</li>                    <li>All the collections</li>                    <li>All the identifiers</li>                    <li>All the formats</li>                    <li>All the logs</li>                  </ul>                </li>       <li>All details regarding the repositories</li>         </ul> <b>" + institution.name + "</b><br /><br /> Are you sure?",
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
        window.location = "/admin/institutions";
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        "url": '/admin/institutions/' + institution.id + "/delete",
        "success": successCallback,
        "error": errorCallback
      };
      wisply.executePostAjax(request);
    }
  };

  /**
  * The constructor activates the listeners
  * @memberof InstitutionAdvanceOptionsModule
  * @class GUI
  * @classdesc It encapsulets the GUI functionality
  */
  var GUI = function GUI() {
  };
  /**
  * It activates all the listeners for the actions
  */
  GUI.activateActionListeners = function() {
    $(".institutions-init-harvest").click(function(event){
      event.preventDefault();
      var object,
      id,
      form,
      xsrf;
      object = $(this);
      id = object.data("id");
      xsrf = wisply.getXSRF();
      form = '<form action="/admin/harvest/init/' + id + '" method="POST">' +
      '<input type="hidden" name="_xsrf" value="' + xsrf + '">' +
      '</form>';
      $(form).submit();
    });
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
  var module = new InstitutionAdvanceOptionsModule();
  wisply.loadModule("admin-institutions-institution-advance-options", module);
});
