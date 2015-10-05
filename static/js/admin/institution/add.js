/* global $, wisply */

/**
* @file Encapsulates the functionality for institutions
* @author Cristian Sima
*/

/**
* @namespace FunctionalityInstitution
*/
var FunctionalityInstitution = function () {
  'use strict';

  /**
  * The constructor activates the listeners
  * @memberof AddInstitution
  * @class Manager
  * @classdesc It encapsulets the functionality for adding an institution
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
      $("#institution-name").focus();
      $("#show-wiki-source").click(this.showWikiSource);
    },
    showWikiSource: function(){
      $("#wiki-source-div").show("fast");
      $("#show-wiki-source").hide("fast");
      $("#institution-wikiURL").focus();
    }
  };

  return {
    Manager: Manager,
  };
};
$(document).ready(function() {
  "use strict";
  var module = new FunctionalityInstitution();
  wisply.functionalityInstitutionModule = module;
  wisply.functionalityInstitutionModule = new module.Manager();
  wisply.functionalityInstitutionModule.init();
});
