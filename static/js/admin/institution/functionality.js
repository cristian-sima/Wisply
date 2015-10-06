/* global $, wisply, server */

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
    this.nameTyper = new wisply.typerModule.Typer("institution-name", function(){
      instance.getWikiByName();
    });
    this.wikier = new wisply.wikierModule.Wikier();
    this.original = server.original;
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
      description = $("#institution-description"),
      logo = $("#institution-logoURL");
      $("#show-wiki-source").click(this.showWikiSource);
      description.elastic();
      logo.on("change keyup paste", function() {
        instance.checkForChanges();
      });
      description.on("change keyup paste", function() {
        instance.checkForChanges();
      });
      $(".discard-description-changes").click(function(){
          $(".description-modified").hide("fast");
          instance.getWikiByName();
      });
      $("#button-get-wiki-by-address").click(function() {
          var url = $("#institution-logoURL").val();
          this.original.wikiReceive = false;
          instance.changeLogo(url);
      });
      $("#button-institution-wikiURL").click(function(){
        instance.getWikiByPage();
      });
    },
    checkForChanges: function(){
        var wikiID = "";
        if(this.original.wikiReceive &&
          this.original.description !== $("#institution-description").val() ||
          this.original.logoURL !== $("#institution-logoURL").val()) {
          wikiID = "NULL";
          this.fired_descriptionModified();
          $("#button-get-wiki-by-address").prop("disabled", false);
        } else {
          wikiID = this.original.wikiID;
          $(".description-modified").hide("fast");
          $("#button-get-wiki-by-address").prop("disabled", true);
        }
        $("#institution-wikiID").val(wikiID);
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
    prepareForWiki: function () {
        this.original.wikiReceive = true;
      $("#institution-logo").html(wisply.getLoadingImage("medium"));

      var name = $("#institution-name"),
        description = $("#institution-description"),
        url = $("#institution-wikiURL"),
        logo = $("#institution-logoURL"),
        institutionURL = $("#institution-institution-URL"),
        submit = $("#institution-submit-button"),
        button = $("#button-institution-wikiURL");

      name.prop( "disabled", true );
      description.prop( "disabled", true );
      url.prop( "disabled", true );
      logo.prop( "disabled", true );
      institutionURL.prop( "disabled", true );
      submit.prop( "disabled", true );
      button.prop( "disabled", true );

      description.val("");
      logo.val("");

    },
    wikiIsDone: function() {
        $("#institution-name").prop( "disabled", false );
        $("#institution-description").prop( "disabled", false );
        $("#institution-wikiURL").prop( "disabled", false );
        $("#institution-logoURL").prop( "disabled", false );
        $("#institution-institution-URL").prop( "disabled", false );
        $("#institution-submit-button").prop( "disabled", false );
        $("#button-institution-wikiURL").prop( "disabled", false );
    },
    getWikiByName: function() {
      var instance = this,
        name = $("#institution-name"),
        newTitle = name.val(),
        descriptionElement = $("#institution-description");
      this.prepareForWiki();
      this.wikier.getByTitle(newTitle, function(err, page){
          instance.wikiIsDone();
          if(err) {
            instance.wikiHasError();
          } else {
            instance.setWikiElements(page);
          }
          name.focus();
        });
    },
    getWikiByPage: function() {
      var instance = this,
        wikiURL = $("#institution-wikiURL"),
        rawPageValue = wikiURL.val(),
        goodPage = rawPageValue.substr(rawPageValue.lastIndexOf('/') + 1);
      this.prepareForWiki();
      this.wikier.getByTitle(goodPage, function(err, page){
          instance.wikiIsDone();
          if(err) {
            instance.wikiHasError();
          } else {
            instance.setWikiElements(page);
          }
          wikiURL.focus();
        });
    },
    wikiHasError: function() {
        this.original.wikiReceive = false;
        this.setDefaultLogo();
    },
    /**
     * It changes the logo, description, urls and wiki page ID to the ones received
     * @param  {object} page The object which contains the elements
     */
    setWikiElements: function(page) {
      var instance = this;
        this.changeLogo(page.thumbnail.source);
        this.changeDescription(page.extract);
        this.setWikiURL(page.fullurl);
        this.setWikiID(page.pageid);
        setTimeout(function(){
          instance.checkForChanges();
        }, 500);
    },
    /**
     * It sets the ID of the wiki page
     * @param {number} newID The new ID
     */
    setWikiID: function (newID) {
        this.original.wikiID = newID;
        $("#institution-wikiID").val(newID);
    },
    /**
     * It changes the wiki URL
     * @param {string} newURL The new URL
     */
    setWikiURL: function (newURL) {
      $("#institution-wikiURL").val(newURL);
    },
    setDefaultLogo: function() {
        this._setLogo("");
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
          this.original.description = text;
    },
    fired_descriptionModified: function() {
        $(".description-modified").show("slow");
    },
    changeLogo: function(picture) {
      var instance = this,
      html = "";
      instance.checkURL(picture, function(isOk) {
        if(!isOk) {
          instance.setDefaultLogo();
        } else {
          instance._setLogo(picture);
        }
      });
    },
    _setLogo: function(url) {
        var logo = "<img class='thumbnail' src='" + url + "' />";
        $("#institution-logoURL").val(url);
        this.original.logoURL = url;
        if(url === "") {
          logo = '<span class="institution-logo glyphicon glyphicon-education institution-logo"></span>';
        }
        $("#institution-logo").html(logo);
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
