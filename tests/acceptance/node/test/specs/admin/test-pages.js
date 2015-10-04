var admin = require("./info.js"),
  user = require("../general/user.js");

accountsPage = {
  title: 'Admin - Accounts'
};

describe('Admin - Test pages', function() {
  admin.init(browser);
  describe("institutions", function() {
    it('list institutions', function(done) {
        browser
        .url('/admin/institutions')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Admin - Institutions");
        })
        .call(done);
    });
    it('add', function(done) {
        browser
        .url('/admin/institutions/add')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Add Institution");
        })
        .call(done);
    });
  });
  describe("repositories", function() {
    it('list repositories', function(done) {
        browser
        .url('/admin/repositories')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe("Admin - Repositories");
        })
        .call(done);
    });
    it('add', function(done) {
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
