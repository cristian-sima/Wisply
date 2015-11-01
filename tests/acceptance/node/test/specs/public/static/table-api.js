
describe('API & Developers', function() {
      it('goes to API & Developers page', function(done){
        browser.
          url('/developer')
          .pause(1000)
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe("API & Developers");
          })
          .call(done);
      });
      it('goes to API page for tables', function(done){
        browser.
            url('/developer/table/list')
            .pause(1000)
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe("API & Developers");
            })
            .call(done);
      });
      var i, tableName,
        downloadTable = function(tableName) {
          var old, newPage;
          it('downloads the table ' + tableName, function(done){
          browser.
              pause(1000)
              .click("#" + tableName)
              .pause(2000)
              .getTitle(function(err, title) {
                  expect(err).toBe(undefined);
                  expect(title).toBe("API & Developers");
              })
              .pause(1500)
              .call(done);
        });
      };
      var tables = ["identifier", "institution", "task", "operation", "process", "repository"];
      for(i=0; i < tables.length; i++) {
        tableName = "download-" + tables[i] + "-table";
        downloadTable(tableName);
      }
    });
