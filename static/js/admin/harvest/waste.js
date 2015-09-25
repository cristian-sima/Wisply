/* global $, wisply,window, data, server*/
/**
 * @file Encapsulates the functionality for managing repositories
 * @author Cristian Sima
 */
/**
 * @namespace Harvest
 */
var InitializationHarvest = function () {
    'use strict';
    /**
     * Creates an empty history
     * @memberof Harvest
     * @class History
     * @classdesc It holds a history of events
     */
    var History = function History(callbackUpdate) {
        this.data = [];
        this.callbackUpdate = callbackUpdate;
    };
    History.prototype =
        /** @lends Harvest.History */
        {
            /**
             * It adds a log message
             * @param [string] content The content of the event
             */
            log: function (content) {
                this.add(content, "LOG");
            },
            /**
             * It adds an error message
             * @param [string] content The content of the message
             */
            logError: function (content) {
                this.add(content, "ERROR");
            },
            /**
             * It adds a warning message
             * @param [string] content The content of the message
             */
            logWarning: function (content) {
                this.add(content, "WARN");
            },
            /**
             * It is called by log, logError and logWarning. It adds an event to history. It assigns the datestamp. It calls the updater
             * @private
             * @param {string} content The content of the message
             * @param {string} type    The type of the message. It can be "LOG", "ERROR" or "WARN"
             */
            add: function (content, type) {
                var datetime;
                /**
                 * It returns the date in a human readable form
                 * @return {string} The date in a human readable form
                 */
                function getHumanDate() {
                    var currentdate = new Date(),
                        datetime = currentdate.getDate() + "." + (currentdate.getMonth() + 1) + "." + currentdate.getFullYear() + " at " + currentdate.getHours() + ":" + currentdate.getMinutes() + ":" + currentdate.getSeconds();
                    return datetime;
                }
                datetime = getHumanDate();
                this.data.unshift({
                    date: datetime,
                    content: content,
                    type: type
                });
                if (this.callbackUpdate) {
                    this.callbackUpdate();
                }
            },
            /**
             * It returns the history in HTML format
             * @return [string] The history in HTML format
             */
            getHTML: function () {
                /**
                 * It creates the HTML header of the table
                 * @return {string} The HTMl header of the table
                 */
                function getHeader() {
                    var header = "";
                    header = "<thead><tr><th class='text-center'>Date</th><th class='text-center'>Category</th><th class='text-center'>Content</th></tr></thead>";
                    return header;
                }
                /**
                 * It creates the body of the table
                 * @param  {string} arrray The events
                 * @return {string}        The HTML body of the table
                 */
                function getBody(arrray) {
                    var result = "<tbody>",
                        i, currentEvent;
                    /**
                     * It returns the type of HTML code
                     * @param  {string} type It can be "LOG", "ERROR" or "WARN"
                     * @return {string}      The HTML code for the type of the event
                     */
                    function getType(type) {
                        var textClass = "",
                            content = "";
                        switch (type) {
                        case "LOG":
                            textClass = "";
                            content = "Event";
                            break;
                        case "ERROR":
                            textClass = "text-danger";
                            content = "Error";
                            break;
                        case "WARN":
                            textClass = "text-warning";
                            content = "Warning";
                            break;
                        }
                        return "<span class='" + textClass + "'>" + content + "</span>";
                    }
                    for (i = 0; i < arrray.length; i++) {
                        currentEvent = arrray[i];
                        result += "<tr>";
                        result += "<td>" + currentEvent.date + "</td>";
                        result += "<td>" + getType(currentEvent.type) + "</td>";
                        result += "<td>" + currentEvent.content + "</td>";
                        result += "</tr>";
                    }
                    result += "</tbody>";
                    return result;
                }
                var html = "<table class='table table=condensed table-hover ''>",
                    events = "";
                html += getHeader();
                events = getBody(this.data);
                html += events;
                html += "</table>";
                return html;
            }
        };

        /**
         * The constructor creates the history, page and stage manager
         * @memberof Harvest
         * @class Manager
         * @classdesc It contains references to the Page object, StageManager and History
         * @param [Repository] repository A reference to the repository
         */
        var HarvestConnection = function HarvestConnection(manager) {
          var host = info.host + "/admin/harvest/init/ws";
          this.manager = manager;
          this.connection = new websockets.Connection(host, this);
        };
        HarvestConnection.prototype =
            /** @lends Harvest.Manager */
            {
        onOpen: function () {
            this.manager.fire_connectionOpen();
        },
        onClose: function () {
            this.manager.fire_connectionClosed();
        },
        onError: function () {
            this.manager.fire_connectionClosed();
        },
        onMessage: function (evt) {
          var message = JSON.parse(evt.data),
            description = "";

          /**
           * It processes the messages received from the server
           * @param [event] evt The event which has been generated
           */
          function createHumanDescription(msg) {
              /**
               * It returns a description of the content of the message
               * @param  {object} content The content of the message
               * @return {string} The description of the content
               */
              function getContentMessage(content) {
                  if (content) {
                      return " with content [" + content + "]";
                  }
                  return ", which does not has content";
              }

              /**
               * It returns a human-readable message for the id of the repository
               * @param  {number} id The id of the repository
               * @return {string} Human readable message
               */
              function getRepo(current, id) {
                  if (current === id) {
                    return "this repository";
                  }
                  return "the repository number " + id;
              }
              return ("I received from server the socket <b>" + msg.Name + "</b>" + getContentMessage(msg.Content) + " for " + getRepo(msg.Repository.id, msg.Repository)+ ".");
          }
          description = createHumanDescription(message);
          history.log(description);
          this.manager.decideAction(message);
        },
        sendMessage: function (msg) {
            this.connection.send(msg);
        },
      };

    /**
     * The constructor creates the history, page and stage manager
     * @memberof Harvest
     * @class Manager
     * @classdesc It contains references to the Page object, StageManager and History
     * @param [Repository] repository A reference to the repository
     */
    var Manager = function Manager(repository) {
        this.history = new History();
        this.page = new Page();
        //  this.stageManager = new StageManager(this);
    };
    Manager.prototype =
        /** @lends Harvest.Manager */
        {

          /**
           *
               var msg = {
                 Name: name,
                 Value: value,
                 Repository: this.manager.repository.id
               };
           */

            /**
             * It activates the listeners
             */
            init: function (serverConfiguration) {
                this.connection = new HarvestConnection(this);
            },
            /**
             * It chooses the action based on name
             * @param  {string} name    The name of the message
             * @param  {string} content The content of the message
             * @param  {number} id The id of repository
             */
            decideAction: function (message) {
                if (message.Repository === this.repository.id) {
                    switch (message.Name) {
                    case "FinishIdentify":
                    case "FinishTestingURL":
                        this.stageManager.currentStage.result(this.stageManager, content);
                        break;
                    case "RepositoryBaseURLChanged":
                        wisply.harvest.repository.url = content;
                        wisply.harvest.restart(2);
                        break;
                    case "RepositoryChangedStatus":
                        wisply.harvest.repository.status = content.NewStatus;
                        wisply.harvest.page.updateRepositoryStatus();
                        switch(content.NewStatus) {
                          case "initializing": {
                              this.stageManager.currentStage.result(this.stageManager);
                          }
                        }
                        break;
                    case "Statistics":
                        this.stageManager.currentStage.update(content.Data);
                        break;
                    case "FinishStage":
                        this.stageManager.currentStage.end(this.stageManager);
                        break;
                    default:
                        this.history.log("This websocket is for the current repository, but it was ignored. Event name <strong>" + name + "</strong> with the content <strong>" + content + "</strong>.");
                    }
                } else {
                    this.history.log("This websocket is not for the current repository. Event name " + name);
                    console.log(content);
                }
            },
            /**
             * It stops the entire process
             */
            stop: function () {
                this.status = "stop";
                this.history.log("The progress has been stopped!");
                this.stageManager.stop();
                this.page.update();
            },
            /**
             * It pauses the entire process
             */
            pause: function () {
                this.status = "=pause";
                this.history.log("The progress has been paused!");
                this.stageManager.pause();
                this.page.warningOccured();
                this.page.update();
            },
            /**
             * It starts the process from a certain stage
             * @param  {number} stageNumber The number of the stage from which it starts again
             */
            restart: function (stageNumber) {
                this.status = "run";
                this.history.log("The progress has been restarted!");
                this.stageManager.restart(stageNumber);
                this.page.restart();
            }
        };
    /**
     *
     * @memberof Harvest
     * @class Page
     * @classdesc It manages the GUI of the page
     */
    var Page = function Page() {
        this.init();
        this.currentTab = "current";
    };
    Page.prototype =
        /** @lends Harvest.Page */
        {
            /**
             * It activates the listeners
             */
            init: function () {
                var instance = this;
                $("#historyButton").click(function () {
                    instance.showHistory();
                });
                $("#modifyButton").click(function () {
                    wisply.harvest.connection.sendMessage("changeRepositoryURL", $("#Source-URL").val());
                    $('#modifyButton').prop('disabled', true);
                });
            },
            /**
             * It change the tab to the history
             */
            showHistory: function () {
                $("#history").html(wisply.getLoadingImage("medium"));
                this.changeTab("history");
            },
            /**
             * It changes the tab to a certian one
             * @param  {string} newTab It can be "history", "current"
             */
            changeTab: function (newTab) {
                this.currentTab = newTab;
                this.update();
            },
            /**
             * It generates the HTML history and assigns it to the tab
             */
            refreshHistory: function () {
                var html = wisply.harvest.history.getHTML();
                $("#history").html(html);
            },
            /**
             * It updates the current view
             */
            update: function () {
                this.updateStages();
                this.updateProcessStatus();
                this.updateRepositoryStatus();
                this.updateHistory();
            },
            /**
             * It updates the history tag
             * @return {[type]}
             */
            updateHistory: function () {
                this.updateHistoryNumber();
                this.updateCurrentTab();
            },
            updateCurrentTab: function() {
                switch (this.currentTab) {
                case "current":
                    break;
                case "history":
                    this.refreshHistory();
                    break;
                }
            },
            /**
             * It updates the number of events in history
             */
            updateHistoryNumber: function () {
                var nr = wisply.harvest.history.data.length;
                $("#historyButton").find("#historyBadge").html(nr);
            },
            /**
             * It updates the list of current stages
             */
            updateStages: function () {
                var manager = wisply.harvest.stageManager,
                    current = manager.current,
                    stages = manager.data,
                    stage, id, html = "",
                    item = "";
                html += "";
                for (id = 0; id < stages.length; id++) {
                    stage = stages[id];
                    if (id === current) {
                        item = '<li class="list-group-item active">' + stage.name + '  <br />';
                        if (stage.showBar) {
                            item += "<div class='progress'><div class='progress-bar progress-bar-success' style='width: 40%''></div></div>";
                        }
                        item += "</li>";
                    } else {
                        if (id < current) {
                            item = '<li class="list-group-item text-muted"><del>' + stage.name + '</del></li>';
                        } else {
                            item = '<li class="list-group-item">' + stage.name;
                        }
                    }
                    html += item;
                }
                if (current === stages.length) {
                    html += '<div class="panel panel-success">  <div class="panel-heading">    <h3 class="panel-title">Done!</h3></div>  <div class="panel-body">    The process is over.  </div></div>';
                    this.processFinished();
                }
                $("#stages").html(html);
                this.updateGeneralIndicator();
            },
            /**
             * It updates the general indicator. This shows the progress based on the number of finished stages
             */
            updateGeneralIndicator: function () {
                var manager = wisply.harvest.stageManager,
                    current = manager.current,
                    total = manager.data.length,
                    percent = 0;
                if (current === "None") {
                    percent = 0;
                } else {
                    percent = (current) / total;
                }
                percent = percent * 100;
                if (this.animateGeneralIndicator) {
                    this.animateGeneralIndicator.finish();
                }
                this.animateGeneralIndicator = $("#generalIndicator").find(".progress-bar").animate({
                    "width": percent + "%"
                }, 100);
            },
            /**
             * It is called when the process has finished. It removes the stripes and makes the bar green
             */
            processFinished: function () {
                this.changeIndicator("success");
            },
            /**
             * It changes the design of indicator for a warning situation
             */
            warningOccured: function () {
                this.changeIndicator("warning");
            },
            /**
             * It changes the design of indicator for an error situation
             */
            errorOcurred: function () {
                this.changeIndicator("danger");
            },
            /**
             * It changes the design of the indicator for a certain situation
             * @param  {string} type The situation. It can be "danger" or "warning" or "success"
             */
            changeIndicator: function (type) {
                var general = $("#generalIndicator");
                general.removeClass("progress-striped");
                general.find(".progress-bar").addClass("progress-bar-" + type);
            },
            /**
             * It activates the stripes for the indicator, it removes the warning design and it updates the general indicator
             */
            restart: function () {
                var general = $("#generalIndicator");
                general.addClass("progress-striped");
                general.find(".progress-bar").removeClass("progress-bar-warning");
                this.updateGeneralIndicator();
            },
            /**
             * It updates the general status of the process
             */
            updateProcessStatus: function () {
                var status = wisply.harvest.stageManager.status,
                    html = "Progress: ";
                switch (status) {
                case "stopped":
                    html += '<span class="label label-danger">Stopped</span>';
                    break;
                case "paused":
                    html += '<span class="label label-warning">Paused</span>';
                    break;
                case "finish":
                    html += '<span class="label label-success">Finish</span>';
                    break;
                case "running":
                    html += wisply.getLoadingImage("small");
                    break;
                default:
                    html += "Problem";
                    break;
                }
                $("#process-status").html(html);
            },
            /**
             * It updates the general status of the process
             */
            updateRepositoryStatus: function () {
                var status = wisply.harvest.repository.status,
                    html = "Status: ",
                    label = "";
                switch (status) {
                case "unverified":
                  label = "info";
                    break;
                case "problems":
                case "verification-failed":
                  label = "danger";
                    break;
                case "upgrading":
                case "verifying":
                case "initializing":
                  label = "warning";
                    break;
                case "ok":
                case "verified":
                  label = "success";
                    break;
                default:
                  console.log("problem");
                  break;
                }
                html += '<span class="label label-' + label + '">' + status + '</span>';
                $("#repository-status").html(html);
            }
        };
    return {
        Manager: Manager,
    };
};
$(document).ready(function () {
    "use strict";
    var harvestModule,
        repositoryModule,
        repository;
    harvestModule = new Harvest();
    repositoryModule = new Repositories();
    repository = new repositoryModule.Repository(data);
    wisply.harvest = new harvestModule.InitManager(repository);
    wisply.harvest.init(server);
});
