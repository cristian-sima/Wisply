var admin = require("./info.js"),
  user = require("../general/user.js");

accountsPage = {
  title: 'Admin - Accounts'
};

describe('Administration pages', function() {
  admin.init(browser);
  describe("institutions", function() {
    it('lists the institutions', function(done) {
        browser
        .url('/admin/institutions')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Admin - Institutions");
        })
        .call(done);
    });
    it('goes to the page for inserting an institution', function(done) {
        browser
        .url('/admin/institutions/add')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Add Institution");
        })
        .call(done);
    });
  });
  describe("Repository", function() {
    it('goes to the page which lists the repositories', function(done) {
        browser
        .url('/admin/repositories')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Admin - Repositories");
        })
        .call(done);
    });
    it('goes to the page which adds a repository', function(done) {
        browser
        .url('/admin/repositories/add')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Add Repository");
        })
        .call(done);
    });
  });
});
