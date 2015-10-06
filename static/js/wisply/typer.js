/* global $, wisply */

/**
* @file Encapsulates the functionality for a typer.
* @author Cristian Sima
*/

/**
* Encapsulates the functionality for the typer.
* @namespace TyperModule
*/
var TyperModule = function () {
  'use strict';

  /**
  * Inits the listenrs. The default interval is 1.5 seconds
  * @memberof TyperModule
  * @class Typer
  * @classdesc It detects when the user has stopped typing and calls a callback
  * @param [string] id The id of the input which will be used to type
  * @param [function] callback The function which is called when the user has finished to type
  * @see http://stackoverflow.com/questions/4220126/run-javascript-function-when-user-finishes-typing-instead-of-on-key-up
  */
  var Typer = function Typer(id, callback) {
      this.timer = {};
      this.interval = 1000;
      this.element = $("#" + id);
      this.callback = callback;
      this.init();
  };
  Typer.prototype =
  /** @lends TyperModule.Typer */
  {
    /**
    * It activates the listeners
    */
    init: function () {
      var instance = this;
      this.element.on('keyup', function () {
        instance.deleteTimer();
        instance.timer = setTimeout(instance.callback, instance.interval);
      });
      this.element.on('keydown', function () {
        clearTimeout(instance.timer);
      });
    },
    /**
     * It deletes the timer
     */
    deleteTimer: function () {
      clearTimeout(this.timer);
    },
    /**
     * It changes the default interval to a custom one
     * @param {number} customInterval The custom interval
     */
    setCustomInterval: function (customInterval) {
        this.interval = customInterval;
    }
  };
  return {
      Typer : Typer,
  };
};

$(document).ready(function() {
  "use strict";
  var module = new TyperModule();
  wisply.typerModule = module;
});
