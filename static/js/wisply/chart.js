/* global $, wisply */

/**
* @file Encapsulates the functionality for a chart.
* @author Cristian Sima
*/

/**
* Encapsulates the functionality for the chart.
* @namespace ChartModule
*/
var ChartModule = function () {
  'use strict';

  /**
  * Inits the listeners.
  * @memberof ChartModule
  * @class Manager
  * @classdesc Manages the charts
  */
  var Manager = function Manager(id, callback) {
    this.showColors = true;
    this.transparent = false;
    this.colors = {};
  };
  Manager.prototype =
  /** @lends ChartModule.Manager */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      var instance = this;
      $("#showColors").click(function(event){
        event.preventDefault();
        var element = $(this);
        if(!instance.showColors) {
          instance.colorWords();
          element.html("<span class='glyphicon glyphicon-text-background text-primary'></span> Remove colors");
        } else {
          $(".word-occurence").css({
            "background-color" : "",
            "color": "",
          });
          element.html("<span class='glyphicon glyphicon-text-background text-info'></span> Show colors");
        }
        instance.showColors = !instance.showColors;
      });
      
      $(".chart-doughnut").each(function(){
        var element = $(this),
        id = element.data("id"),
        ctx = element[0].getContext("2d"),
        data = instance.getData(id);
        new Chart(ctx).Doughnut(data);
        instance.colorWords();
      });
      $(".chart-radar").each(function(){
        var element = $(this),
        id = element.data("id"),
        ctx = element[0].getContext("2d"),
        data = instance.getData(id);
        new Chart(ctx).PolarArea(data);
        instance.colorWords();
      });
      $(".chart-pie").each(function(){
        var element = $(this),
        id = element.data("id"),
        ctx = element[0].getContext("2d"),
        data = instance.getData(id);
        new Chart(ctx).Pie(data);
        instance.colorWords();
      });
    },
    colorWords: function() {
      var instance = this;
      $(".word-occurence").each(function(){
        var element = $(this),
        word = element.data("word"),
        color = instance.getColorForWord(word);
        element.css({"background-color": color.background,
        "color": color.font });
      });
    },

    /**
    * It gets the data for the chart. Also, it processes it
    * @param  {string} id The ID of the analyse
    * @return {object} The data
    */
    getData : function(id) {
      var data = dataPool[id],
      i, occurence, newSet = [], newItem;
      // transform it
      for(i=0; i<data.length;i++) {
        occurence = data[i];
        newItem = {
          "label": occurence.Word,
          "value": occurence.Counter,
          color : this.getColorForWord(occurence.Word).background,
        };
        newSet.push(newItem);
      }
      return newSet;
    },
    /**
    * It checks to see if the color for the word is already stored. If not, it generates a new one
    * @param  {string} word The word
    * @return {object} The color in RGB format
    */
    getColorForWord: function(word) {
      if(!this.colors[word]) {
        this.colors[word] = this.getRandomColor();
      }
      return this.colors[word];
    },
    /**
    * It returns a random color
    * @return {object} A random color for bg and font
    */
    getRandomColor: function() {
      function getContrastYIQ(hexcolor){
        var r = parseInt(hexcolor.substr(0,2),16);
        var g = parseInt(hexcolor.substr(2,2),16);
        var b = parseInt(hexcolor.substr(4,2),16);
        var yiq = ((r*299)+(g*587)+(b*114))/1000;
        return (yiq >= 128) ? 'black' : 'white';
      }
      var obj = {},
      background = Math.floor(Math.random()*16777215).toString(16);
      // credits http://stackoverflow.com/questions/11070007/style-each-div-with-a-different-color-using-jquery-or-javascript
      obj.background = '#' + background;
      obj.font = getContrastYIQ(background);
      return obj;
    },
  };
  return new Manager();
};

$(document).ready(function() {
  "use strict";
  var module = new ChartModule();
  wisply.loadModule("chart", module);
});
