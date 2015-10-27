/* global $, wisply */

/**
* @file Encapsulates the functionality for managing captcha images
* @author Cristian Sima
*/

/**
* @namespace CaptchaModule
*/
var CaptchaModule = function () {
  'use strict';


  /**
  * The constructor does nothing
  * @class Manager
  * @memberof CaptchaModule
  * @classdesc The Manager for Captcha images
  */
  function Captcha(o) {
    this.ID = o.ID;
    this.element = o.element;
  }
  Captcha.prototype =
  /** @lends Captcha.CaptchaModule */
  {

  };
  return {
    Captcha: Captcha
  };
};

$(document).ready(function() {
  wisply.loadModule("captcha", new CaptchaModule());
});
