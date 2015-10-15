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
    this.resourcePerPage = 15;
    this.delimitator = {
      verb : "*",
      parameter: "-",
      insideVerb: "|",
    };
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
      function initHash() {
        $(window).on('hashchange', function() {
          instance.hashChanged();
        });
      }

      function initKeys() {
        var shortcuts = [{
            "type": "keyup",
            "key": "Ctrl+left",
            "callback": function () {
              wisply.publicRepositoryModule.manager.showPrevious();
            }
          }, {
            "type": "keyup",
            "key": "Ctrl+right",
            "callback": function () {
              wisply.publicRepositoryModule.manager.showNext();
            }
          }];
        wisply.shortcutManager.activate(shortcuts);
      }
      initHash();
      initKeys();
    },
    /**
    * It loads the resources
    */
    getResources: function () {
      var instance = this;
      this.showLoading();
      $.ajax({
        url:"/api/repository/resources/" + this.repository.id + "/get/" + this.min + "/" + this.resourcePerPage,
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
      this.updateGUI();
    },
    activateListeners: function() {
      var instance = this;
      function initButtons() {
        $(".next").click(function(event){
          event.preventDefault();
          if (!$(this).hasClass("disabled")) {
            instance.showNext();
          }
        });
        $(".previous").click(function(event){
          event.preventDefault();
          if (!$(this).hasClass("disabled")) {
            instance.showPrevious();
          }
        });
      }
      initButtons();
    },
    showNext: function(event) {
      var newMin;

      newMin = parseInt(this.min, 10) + this.resourcePerPage;

      this.changeMin(newMin);

      this.goUp();
      this.updateHash();
    },
    showPrevious: function(event) {
      var newMin;

      newMin = parseInt(this.min, 10) - this.resourcePerPage;

      this.changeMin(newMin);

      this.goUp();
      this.updateHash();
    },
    changeMin: function(newValue) {
      var value = 0;
      if ((!isInt(newValue)) || !newValue || newValue === "" || parseInt(newValue, 10) < 0) {
        value = 0;
      } else {
        value = newValue;
      }
      this.min = parseInt(value, 10);
    },
    updateHash: function() {
      var instance = this;

      function getList() {
        return "list" + instance.delimitator.insideVerb + instance.min + instance.delimitator.parameter + instance.resourcePerPage;
      }
      function getVerbs() {
        var verbs = [];
        verbs.push(getList());
        return verbs.join(instance.delimitator.verb);
      }
      window.location.hash = getVerbs();
    },
    hashChanged: function() {
      var h = window.location.hash,
      hash = h.substr(h.indexOf('#')+1),
      verbs,
      instance = this;

      function extractVerbs(URL) {
        var ret= [],
        verbs =  URL.split(instance.delimitator.verb);
        var i, verb, index, extracted;
        for (i=0; i < verbs.length; i++) {
          verb = verbs[i];
          extracted = extractElements(verb);
          ret.push(extracted);
        }
        return ret;
      }

      function extractElements(URLString) {
        var elements = {},
        el;

        el = URLString.split(instance.delimitator.insideVerb);

        elements.verb = el[0];
        if (el[1]) {
          elements.parameters = el[1].split(instance.delimitator.parameter);
        }
        return elements;
      }

      this.verbs = extractVerbs(hash);
      this.updateListVerb();
      this.getResources();
    },
    getVerb: function(name) {
      var verb,
      instance = this,
      verbs = instance.verbs;
      for (verb in verbs) {
          if (verbs.hasOwnProperty(verb)) {
              if (verbs[verb].verb === name) {
                return verbs[verb];
              }
          }
      }
    },
    updateListVerb: function() {
    var verb = this.getVerb("list");
      if(verb) {
        this.changeMin(verb.parameters[0]);
        this.changeResourcesPerPage(verb.parameters[1]);
      }
    },
    changeResourcesPerPage: function (newValue) {
        this.resourcePerPage = parseInt(newValue, 10);
    },
    updateGUI: function() {
      var instance = this;

      function updatePreviousButton() {
        if (instance.min < instance.resourcePerPage) {
          $(".previous").addClass("disabled");
        } else {
          $(".previous").removeClass("disabled");
        }
      }

      function updateNextButton() {
        if (instance.min + instance.resourcePerPage >= parseInt(instance.repository.totalRecords, 10)) {
          $(".next").addClass("disabled");
        } else {
          $(".next").removeClass("disabled");
        }
      }

      function updateButtons() {
        updatePreviousButton();
        updateNextButton();
      }
      updateButtons();
    },
    goUp: function() {
      var listPosition = parseInt($("#repository-before-resources").offset().top, 10) - 70,
      currentPosition = parseInt($(document).scrollTop(), 10);

      function isAround(current, point) {
        var range = 50;
        return (current >= point - range && current <= point) ||
              (current >= point && current <= point + range);
      }
      console.log(listPosition + " - " + currentPosition);
      // if (!isAround(currentPosition, listPosition)) {
        $('html, body').animate({
          scrollTop: listPosition,
        }, 100);
      // }
    },
    changeResources: function (html) {
      $("#repository-resources").html(html);
    },
    updateTop: function(html) {
        $("#repository-top").html(html);
    },
    updateBottom: function(html) {
        $("#repository-bottom").html(html);
    }
  };
  function isInt(value) {
    return !isNaN(value) &&
         parseInt(Number(value)) == value &&
         !isNaN(parseInt(value, 10));
  }
  return {
    Manager: Manager
  };
};
$(document).ready(function() {
  "use strict";
  var module = new PublicRepository();
  wisply.publicRepositoryModule = module;
  wisply.publicRepositoryModule.manager = new module.Manager(server.repository);
  wisply.publicRepositoryModule.manager.init();
});
