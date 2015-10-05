/* global $, wisply */

/**
* It contains the object typer
* @global
* @see Wisply.App
*/

/**
* @file Encapsulates the functionality for a typer.
* @author Cristian Sima
*/

/**
* Inits the listenrs
* @memberof ModifyInstitution
* @class Manager
* @classdesc Encapsulates the functionality for typing
* @param [string] id The id of the input which will be used to type
* @param [function] callback The function which is called when the user has finished to type
* @see http://stackoverflow.com/questions/4220126/run-javascript-function-when-user-finishes-typing-instead-of-on-key-up
*/
var Typer = function Typer(id, callback) {
    this.timer = {};
    this.interval = 3000;
    this.element = $(id);
    this.init();
    this.callback = callback;
};
Typer.prototype =
/** @lends Institutions.Manager */
{
  /**
  * It activates the listeners
  */
  init: function () {
    var instance = this;
    $input.on('keyup', function () {
      intance.clear();
      intance.timer = setTimeout(instance.callback, instance.interval);
    });
    $input.on('keydown', function () {
      clearTimeout(instance.timer);
    });
  },
  /**
   * It clears the typing timer
   */
  clear: function () {
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
$(document).ready(function() {
  "use strict";
  var module = new TyperModule();
  wisply.typerModule = module;
});
