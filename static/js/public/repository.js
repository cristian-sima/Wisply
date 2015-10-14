/* global $, wisply, server */

/**
* @file Encapsulates the functionality for displaying repositories
* @author Cristian Sima
*/


/**
* @namespace PublicRepository
*/
var PublicRepository = function () {
  'use strict';

  /**
  * The constructor activates the listeners
  * @memberof PublicRepository
  * @class Manager
  * @classdesc It encapsulets the functionality for the repositories
  * @param {object} currentRepository Represents the current repository
  */
  var Manager = function Manager(currentRepository) {
    this.repository = currentRepository;
    this.min = 0;
    this.max = 15;
    this.resourcePerPage = 15;
  };
  Manager.prototype =
  /** @lends PublicRepository.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      this.initListeners();
      this.hashChanged();
    },
    initListeners: function() {
      var instance = this;
      $(window).on('hashchange', function() {
        instance.hashChanged();
      });
    },
    /**
    * It loads the resources
    */
    getResources: function () {
      var instance = this;
      this.showLoading();
      $.ajax({
        url:"/api/repository/resources/" + this.repository.id + "/get/" + this.min + "/" + this.max,
        success: function(html) {
          instance.fired_NewResources(html);
        }
      });
    },
    showLoading: function() {
      $("#repository-resources").html('<div class="text-center">' + wisply.getLoadingImage("medium") + '</div>');
    },
    fired_NewResources: function(html) {
        this.changeResources(html);
        this.activateListeners();
    },
    activateListeners: function() {
      var instance = this;
        $(".next").click(function(event){
          event.preventDefault();
          instance.showNext();
        });
        $(".previous").click(function(event){
          event.preventDefault();
          instance.showPrevious();
        });
    },
    showNext: function(event) {
      var min, max;
      min = this.max;
      max = parseInt(this.max, 10) + this.resourcePerPage;
      window.location.hash = min + "-" + max;
      this.goUp();
    },
    showPrevious: function(event) {
      var min, max;
      min = parseInt(this.min, 10) - this.resourcePerPage;
      if (min < 0) {
        min = 0;
      }
      max = parseInt(this.max, 10) - this.resourcePerPage;
      window.location.hash = min + "-" + max;
      this.goUp();
    },
    goUp: function() {
      $('html, body').animate({
        scrollTop: 0,
      }, 100);
    },
    changeResources: function (html) {
      $("#repository-resources").html(html);
    },
    hashChanged: function() {
        var h = window.location.hash,
          hash = h.substr(h.indexOf('#')+1),
        elements = hash.split("-");
        console.log(elements);
        if (!elements[0] || elements[0] === "" || parseInt(elements[0], 10) < 0) {
          this.min = 0;
        } else {
          this.min = elements[0];
        }
        if (!elements[1] || elements[1] === "" || parseInt(elements[1], 10) <= parseInt(this.min)) {
          this.max = this.resourcePerPage;
        } else {
          this.max = elements[1];
        }
        this.getResources();
    }
  };
  return {
    Manager: Manager
  };
};
$(document).ready(function() {
  "use strict";
  var module = new PublicRepository();
  wisply.publicRepositoryModule = module;
  wisply.publicRepositoryModule = new module.Manager(server.repository);
  wisply.publicRepositoryModule.init();
});
