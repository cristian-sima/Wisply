
var  user = {
      name: "Jameson Henry",
      email: "henry@oxford.ac.uk",
      password: "my-strong-password"
    },
    settings = {
      separator: "::"
    },
     connectionCookie = {},
   hackedCookie = "";

// creates an ccount
describe('API - Tables', function() {
      it('go to api page', function(done){
        browser.
          url('/api')
          .pause(1000)
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe("API & Developers");
          })
          .call(done);
      });
      it('go to tables tapage', function(done){
        browser.
            url('/api/table/list')
            .pause(1000)
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe("API & Developers");
            })
            .call(done);
      });
      var i, tableName,
        downloadTable = function(tableName) {
          it('download the table ' + tableName, function(done){
          browser.
              pause(1000)
              .click("#" + tableName)
              .pause(2000)
              .getTitle(function(err, title) {
                  expect(err).toBe(undefined);
                  expect(title).toBe("API & Developers");
              })
              .keys("U+E00C")
              .pause(1000)
              .call(done);
        });
      };
      var tables = ["identifier", "institution", "task", "operation", "process", "repository"];
      for(i=0; i < tables.length; i++) {
        tableName = "download-" + tables[i] + "-table";
        downloadTable(tableName);
      }
    });
