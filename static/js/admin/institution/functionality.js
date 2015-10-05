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
    this.description = "";
  };
  Manager.prototype =
  /** @lends FunctionalityInstitution.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.activateListeners();
      wisply.activateTooltip();
      $("#institution-name").focus();
    },
    /**
    * It activates the listener for deleting a institution
    * @fires InstitutionsManager#confirmDelete
    */
    activateListeners: function () {
      var instance = this,
      description = $("#institution-description");
      $("#show-wiki-source").click(this.showWikiSource);
      description.elastic();
      description.on("change keyup paste", function() {
        if(instance.description !== $("#institution-description").val()) {
          instance.fired_descriptionModified();
        } else {
          $(".description-modified").hide("fast");
        }
      });
      $("#discard-description-changes").click(function(){
          $(".description-modified").hide("fast");
          instance.update();
      });
    },
    /**
     * It shows the field for wiki source and hides the link
     */
    showWikiSource: function(event){
      event.preventDefault();
      $("#wiki-source-div").show("fast");
      $("#show-wiki-source").hide("fast");
      $("#institution-wikiURL").focus();
    },
    /**
     * It is called when the name of the institution is changed.
     */
    fired_nameChanged: function () {
        this.update();
    },
    get: function () {
        // get wiki id and description by name

    },
    update: function() {
      var instance = this,
        newName = $("#institution-name").val(),
        html = "",
        descriptionElement = $("#institution-description"),
        callbackPicture;
      this.wikier.changeSubject(newName);
      $("#institution-logo").html(wisply.getLoadingImage("medium"));
      this.wikier.getPicture(function(err, page){
        if(err) {
          instance.setDefaultLogo();
        } else {
          var picture = page.thumbnail;
          instance.changeLogo(picture);
        }
      });
      descriptionElement.html("Please wait");
      descriptionElement.prop( "disabled", true );
      this.wikier.getDescription(function(err, description){
        var text = "";
          if(!err) {
            text = description;
          }
          instance.changeDescription(text);
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
           limit = 1000,
           text = "";
        function cutDescription(text) {
          var start = 0,
          endOfParagraph = 0;
          if(text.length > limit) {
          endOfParagraph = newDescription.lastIndexOf("\n", limit);
          text = text.substring(0, endOfParagraph);
          }
          return text;
        }

          text = cutDescription(newDescription);
          description.val(text);
          description.elastic();
          this.description = text;

        $(description).prop( "disabled", false );
    },
    fired_descriptionModified: function() {
        $(".description-modified").show("slow");
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
