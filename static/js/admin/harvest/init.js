/* global $, wisply,window, host*/
/**
 * @file Encapsulates the functionality for managing repositories
 * @author Cristian Sima
 */
/**
 * @namespace Harvest
 */
var Harvest = function () {
    'use strict';
    /**
     * Creates an empty history
     * @memberof Harvest
     * @class History
     * @classdesc It holds a history of events
     */
    var History = function History() {
        this.data = [];
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
                if (wisply.harvest) {
                    wisply.harvest.page.update();
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
     * Starts the connection and inits the listeners
     * @memberof Harvest
     * @class Connection
     * @classdesc It represents a websocket connection
     * @param {function} processor A callback called when a message is received
     */
    var Connection = function Connection(processor) {
        this.value = new WebSocket("ws://" + host + "/admin/harvest/init/ws");
        this.processor = processor;
        this.initListeners();
    };
    Connection.prototype =
        /** @lends Harvest.Connection */
        {
            /**
             * It
             */
            initListeners: function () {
                this.value.onopen = function () {
                    wisply.harvest.history.log("The websocket connection has been created");
                    $("#connectionStatus").html("<span class='text-success'>WebSocket connection established</span>");
                };
                this.value.onclose = function () {
                    wisply.harvest.history.logError("The webscoket connection is closed");
                    $("#connectionStatus").html("<span class='text-danger'>No WebSocket connection</span>");
                    wisply.harvest.stop();
                };
                this.value.onmessage = this.processor;
                this.value.onerror = function () {
                    wisply.harvest.history.logWarning("There was a an error with web scoket connection");
                };
            },
            /**
             * It sends a message
             * @param  {string} name The name of the message
             * @param  {object} value The value of the message
             */
            sendMessage: function (name, value) {
                var msg = {
                    name: name,
                    value: value
                };
                this.value.send(JSON.stringify(msg));
            }
        };
    /**
     * Saves the stages
     * @memberof Harvest
     * @class StageManager
     * @classdesc It encapsulets the functionality for the sources
     * @param [Manager] repositoriesManager The reference to the repositories manager
     */
    var StageManager = function StageManager(repositoriesManager) {
        this.status = "stopped";
        this.repo = repositoriesManager;
        // stages
        this.data = [{
            name: "Prepare resources",
            id: 0,
            /**
             * It prepares the resources
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                this.paint();
                var manager = stageManager;
                setTimeout(function () {
                    manager.firedStageFinished();
                }, 5000);
            },
            /**
             * It shows the loading image
             */
            paint: function () {
                $('#current').html(wisply.getLoadingImage("big"));
            }
        }, {
            name: "Connect to server",
            id: 1,
            /**
             * It tries to connect to server using websockets and saves the connection
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                var manager = stageManager.repo;
                if (window.WebSocket) {
                    manager.connection = new Connection(function (data) {
                        wisply.harvest.processMessage(data);
                    });
                } else {
                    this.complain();
                    return;
                }
                setTimeout(function () {
                    stageManager.firedStageFinished();
                }, 400);
            },
            /**
             * It tells the user the connection was not done
             */
            complain: function () {
                $('#current').html("Wisply was not able to realize the connection. Your browser does not support WebSockets");
            },
            /**
             * It shows the loading image
             */
            paint: function () {
                $('#current').html(wisply.getLoadingImage("big"));
            }
        }, {
            name: "Verify URL address",
            id: 2,
            showBar: true,
            /**
             * It tries to verify the address
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                this.paint();
                this.disableModifyURL();
                var instance = stageManager;
                instance.repo.history.log("Indentifing the source");
                instance.repo.connection.sendMessage("testURL", $("#Source-URL").val());
            },
            /**
             * It checks if the server has verified the URL
             * @param  {Manager} stageManager The reference to the repositories manager
             * @param  {object} content      The content of the message from server
             */
            result: function (stageManager, content) {
                if (content === "true") {
                    stageManager.repo.history.log("The URL is valid");
                    this.end();
                    stageManager.firedStageFinished();
                } else {
                    this.complain(stageManager);
                }
            },
            /**
             * It shows the loading image
             */
            paint: function () {
                $('#current').html(wisply.getLoadingImage("big"));
            },
            /**
             * It tells the user that the URL was not valid. It allows the user to edit it
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            complain: function (stageManager) {
                $('#current').html("The URL is not valid or the address can not be visited. Please correct it and click 'Modify'");
                stageManager.repo.history.log("The URL is not valid!");
                stageManager.repo.pause();
                this.enableModifyURL();
            },
            /**
             * It disables the possibility to modify the URL
             */
            disableModifyURL: function () {
                $('#modifyButton').prop('disabled', true);
                $('#Source-URL').prop('disabled', true);
            },
            /**
              * It enables the possibility to modify the URL
             */
            enableModifyURL: function () {
                $('#modifyButton').prop('disabled', false);
                $('#Source-URL').prop('disabled', false);
            },
            /**
             * It hides the URL and button for modify and shows the name of the repository
             */
            end: function () {
                $("#URL-input").toggle();
                $("#Name-Repository").toggle();
                $("#modifyButton").hide();
            }
        }, {
            name: "Identify Source",
            id: 3,
            /**
             * It tells the server to identify the source
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                var instance = stageManager;
                instance.repo.history.log("Indentifing the source");
                instance.repo.connection.sendMessage("identify", "something");
            },
            /**
             * It checks if the server has identified the repository
             * @param  {object} indentifyInfo The value of the message from the server
             */
            result: function (stageManager, indentifyInfo) {
                if (indentifyInfo.state === true) {
                    this.paint(indentifyInfo.data.Identify);
                    stageManager.repo.history.log("The source has been identified");
                    stageManager.firedStageFinished();
                } else {
                    stageManager.repo.history.log("There has error during identification!");
                    stageManager.repo.errorOcurred();
                }
            },
            /**
             * It shows in the current table the details about the repository
             */
            paint: function (data) {
                /**
                 * It returns the HTML table for an object
                 * @param  {object} data The object
                 * @return {string} HTML table for the object's data
                 */
                function getHTML(data) {
                    var html = "";
                    html += '<ul class="list-group text-left">';
                    for (var property in data) {
                        if (data.hasOwnProperty(property)) {
                            if (property === "Description") {
                                continue;
                            } else if (typeof data[property] === 'object') {
                                html += '<li class="list-group-item"> ' + property;
                                html += getHTML(data[property]);
                                html += '</li>';
                            } else {
                                html += "  <li class='list-group-item'>";
                                html += property + ": <strong>" + data[property] + "</strong>";
                                html += "</li>";
                            }
                        }
                    }
                    html += "</ul>";
                    return html;
                }
                var html = getHTML(data);
                $("#current").html(html);
            }
        }];
        this.current = "None";
        this.stage = {};
    };
    StageManager.prototype =
        /** @lends Harvest.StageManager */
        {
            /**
             * It starts the manager. It calls the first stage
             */
            start: function () {
                this.current = 0;
                this.performStage(0);
            },
            /**
             * It calls the next stage. If there are no stages, it calls firedEnd
             */
            next: function () {
                this.current++;
                if (this.current >= this.data.length) {
                    this.firedEnd();
                } else {
                    this.performStage(this.current);
                }
            },
            /**
             * It performs a stage
             * @param  {number} id The id of the stage
             */
            performStage: function (id) {
                var stage = this.data[id];
                this.status = "running";
                this.repo.history.log("Starting stage <b>" + this.name + "</b>...");
                this.repo.page.update();
                this.stage = stage;
                stage.perform(this);
            },
            /**
             * It forces the manager to stop. It forces the current stage to stop
             */
            forceStop: function () {
                this.repo.history.log("The stage manager has been forced to stop.");
                this.current = "Stopped";
                this.state = "stopped";
                if (this.stage.stop) {
                    this.stage.stop();
                }
            },
            /**
             * Called when a stage has finished. It updates the page and calls the next stage
             */
            firedStageFinished: function () {
                if(this.state === "stopped" || this.state === "paused") {
                this.repo.history.log("Imposible to continue!");
                } else {
                  this.repo.page.update();
                  this.repo.history.log("Stage " + (this.current + 1) + " finished!");
                  this.next();
                }
            },
            /**
             * It is called when all the stages has been called. It updates the page
             */
            firedEnd: function () {
                this.status = "finish";
                this.repo.page.update();
                this.repo.history.log("The process has been finished!");
            },
            /**
             * It performs again a stage
             * @param  {number} number The id of the stage
             */
            restart: function (number) {
                this.repo.history.log("Restarting from stage " + (number + 1) + "...");
                this.current = number - 1;
                this.next();
            },
            /**
             * It pauses the manager
             */
            pause: function () {
                this.status = "paused";
            }
        };
    /**
     * The constructor creates the history, page and stage manager
     * @memberof Harvest
     * @class Manager
     * @classdesc It contains references to the Page object, StageManager and History
     */
    var Manager = function Manager() {
        this.history = new History();
        this.history.log("The manager has started.");
        this.page = new Page();
        this.stageManager = new StageManager(this);
    };
    Manager.prototype =
        /** @lends Harvest.Manager */
        {
            /**
             * It activates the listeners
             */
            init: function () {
                var instance = this;
                this.page.update();
                instance.stageManager.start();
            },
            /**
             * It processes the messages received from the server
             * @param [event] evt The event which has been generated
             */
            processMessage: function (evt) {
                var msg = JSON.parse(evt.data);

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
                this.history.log("I received the socket [<b>" + msg.name + "</b>]" + getContentMessage(msg.content));
                this.chooseAction(msg.name, msg.content);
            },
            /**
             * It choose the action based on name
             * @param  {string} name    The name of the message
             * @param  {string} content The content of the message
             */
            chooseAction: function (name, content) {
                switch (name) {
                case "FinishIdentify":
                case "FinishTestingURL":
                    this.stageManager.stage.result(this.stageManager, content);
                    break;
                }
            },
            /**
             * It stops the entire process
             */
            stop: function () {
                this.stageManager.forceStop();
                this.page.errorOcurred();
                this.page.update();
            },
            /**
             * It pauses the entire process
             */
            pause: function () {
                this.stageManager.pause();
                this.page.warningOccured();
                this.page.update();
            },
            /**
             * It starts the process from a certain stage
             * @param  {number} stageNumber The number of the stage from which it starts again
             */
            restart: function (stageNumber) {
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
                    wisply.harvest.restart(2);
                });
            },
            /**
             * It change the tab to the history
             */
            showHistory: function () {
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
                switch (this.currentTab) {
                case "current":
                    break;
                case "history":
                    this.refreshHistory();
                    break;
                }
                this.updateStages();
                this.updateHistoryNumber();
                this.updateStatus();
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
            updateStatus: function () {
                var status = wisply.harvest.stageManager.status,
                    html = "Status: ";
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
            }
        };
    return {
        Manager: Manager
    };
};
$(document).ready(function () {
    "use strict";
    var module = new Harvest();
    wisply.harvest = new module.Manager();
    wisply.harvest.init();
});
