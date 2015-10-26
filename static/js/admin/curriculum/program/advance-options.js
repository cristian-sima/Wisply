/* global $, wisply */

/**
* @file Encapsulates the functionality for managing the advance options for
* programs
* @author Cristian Sima
*/


/**
* @namespace AdanceOptionsProgram
*/
var AdanceOptionsProgram = function () {
  'use strict';

  /**
  * The constructor does nothing important
  * @class Program
  * @memberof Programs
  * @classdesc It represents a program
  * @param {object} info It contains the information regarding the program (id, name)
  */
  var Program = function Program(info) {
    this.id = info.id;
    this.name = info.name;
  };

  /**
  * The constructor activates the listeners
  * @memberof AdanceOptionsProgram
  * @class Manager
  * @classdesc It encapsulets the functionality for the programs
  */
  var Manager = function Manager() {
  };
  Manager.prototype =
  /** @lends AdanceOptionsProgram.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.activateListeners();
    },
    /**
    * It activates the listener for deleting a program
    * @fires ProgramsManager#confirmDelete
    */
    activateListeners: function () {
      var instance = this;
      $(".deleteProgramButton").click(function(event) {
        event.preventDefault();
        var object,program;
        object = $(this);
        program = new Program({
          id: object.data("id"),
          name: object.data("name"),
        });
        instance.confirmDelete(program);
      });
    },
    /**
    * It is called when the user wants to delete a program. It asks for confirmation
    * @param  {Program} program The reference to the program object
    */
    confirmDelete: function (program) {
      /**
       * It focuses the password input after the promt is shown
       */
      function focusPassword() {
        setTimeout(function(){
          $("#promt-password").focus();
        }, 500);
      }
      var msg = this.getDeleteDialogMessage(program);
      wisply.message.dialog(msg);
      focusPassword();
    },
    /**
    * It returns the object which contain the arguments for the confirmation dialog
    * @param  {string} type The type of action: "delete", "clear"
    * @return {Object}        The arguements for dialog
    * @see http://bootboxjs.com/
    */
    getDeleteDialogMessage: function (program) {
      var buttons,
      cancelButton,
      msg,
      mainButton,
      instance = this,
      programCopy = program;

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
        callback: function() {
          var password, title;
          title = "Removing the program <strong>" + programCopy.name + "</strong>";
          password = $("#promt-password").val();
          wisply.message.tellToWait(title);
          instance.delete(programCopy, password);
        },
      };
      buttons = {
        "cancel": cancelButton,
        "main": mainButton
      };
      msg = {
        title: "Type your password",
        buttons: buttons,
        onEscape: true,
        message: '<input class="bootbox-input bootbox-input-text form-control" autocomplete="off" type="password" id="promt-password" />',
      };
      return msg;
    },
    /**
    * It delets a program
    * @param  {Program} program The program object
    * @param {string} password The password from the user
    */
    delete: function (program, password) {
      var request,
      successCallback,
      errorCallback;

      /**
      * The callback called when the program has been deleted. It shows a message and reloads the page in 2 seconds
      */
      successCallback =  function () {
        //wisply.message.showSuccess("The program has been removed! Refreshing page...");
        window.location = "/admin/curriculum";
      };

      /**
      * The callback called when there was a problem. It shows a message
      */
      errorCallback = function () {
        wisply.message.showError("There was a problem with your request!");
      };

      request = {
        url: '/admin/curriculum/programs/' + program.id + "/delete",
        success: successCallback,
        error: errorCallback,
        data: {
          password: password,
        },
      };
      wisply.executePostAjax(request);
    },
  };

  return {
    Program: Program,
    Manager: Manager,
  };
};
$(document).ready(function() {
  "use strict";
  var module = new AdanceOptionsProgram();
  wisply.advanceOptionsProgramModule = module;
  wisply.advanceOptionsProgramModule.manager = new module.Manager();
  wisply.advanceOptionsProgramModule.manager.init();
});
