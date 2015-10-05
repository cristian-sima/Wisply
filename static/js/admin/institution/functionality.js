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
  * @memberof FunctionalityInstitution
  * @class Manager
  * @classdesc It encapsulets the functionality for adding an institution
  */
  var Manager = function Manager() {
    var instance = this;
    this.typer = new wisply.typerModule.Typer("institution-name", function(){
      instance.fired_nameChanged();
    });
    this.wikier = new wisply.wikierModule.Wikier();
  };
  Manager.prototype =
  /** @lends FunctionalityInstitution.Manager */
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
      $("#institution-description").elastic();
    },
    /**
     * It shows the field for wiki source and hides the link
     */
    showWikiSource: function(){
      $("#wiki-source-div").show("fast");
      $("#show-wiki-source").hide("fast");
      $("#institution-wikiURL").focus();
    },
    /**
     * It is called when the name of the institution is changed.
     */
    fired_nameChanged: function () {
        var instance = this,
          newName = $("#institution-name").val(),
          html = "",
          callbackPicture;
        this.wikier.changeSubject(newName);
        $("#institution-logo").html(wisply.getLoadingImage("medium"));
        this.wikier.getPicture(function(err, picture){
          if(err) {
            instance.setDefaultLogo();
          } else {
            instance.changeLogo(picture);
          }
        });
        this.wikier.getDescription(function(err, description){
            if(!err) {
              instance.changeDescription(description);
            }
        });
    },
    setDefaultLogo: function() {
        var html = '<span class="institution-logo glyphicon glyphicon-education institution-logo"></span>';
        this._setLogo(html);
    },
    _setLogo: function(logo) {
        $("#institution-logo").html(logo);
    },
    /**
     * It checks if the description is empty. If so it populates it with the description.
     * @param  {string} newDescription The new description
     */
    changeDescription: function(newDescription) {
        var description = $("#institution-description"),
           limit = 1000;
        function cutDescription(text) {
          var start = 0,
          endOfParagraph = 0;
          if(text.length > limit) {
          endOfParagraph = newDescription.lastIndexOf("\n", limit);
          text = text.substring(0, endOfParagraph);
          }
          return text;
        }
        if(description.val() === "") {
          description.val(cutDescription(newDescription));
          description.elastic();
        }
    },
    changeLogo: function(picture) {
      var instance = this,
      html = "";
      instance.checkURL(picture.source, function(isOk) {
        if(!isOk) {
          instance.setDefaultLogo();
        } else {
          var logo = "<img class='thumbnail' src='" + picture.source + "' />";
          instance._setLogo(logo);
        }
      });
    },
    checkURL: function(URL, callback) {
      $.ajax({
          type: 'GET',
          url: URL,
          success: function() {
            callback(true);
          },
          error: function() {
            callback(false);
          }
      });
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
