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
  describe("institutions", function() {
    it('goes to institutions page', function(done) {
        browser
        .url('/admin/institutions')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
    it('goes to institutions page', function(done) {
        browser
        .url('/admin/institutions/add')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
  });
  describe("developer", function() {
    it('goes to developer admin page', function(done) {
        browser
        .url('/admin/developer')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
    it('goes to page which makes the table public', function(done) {
        browser
        .url('/admin/developer/add')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
  });
  describe("log", function() {
    it('goes to log page', function(done) {
        browser
        .url('/admin/log')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
    it('goes to advance page for all', function(done) {
        browser
        .url('/admin/log/advance-options')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
  });
  describe("curricula", function() {
    it('goes to curricula page', function(done) {
        browser
        .url('/admin/curriculum')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
    it('goes to page which adds a new program of study', function(done) {
        browser
        .url('/admin/curriculum/add')
        .isExisting('#full-logo').then(function(isExisting){
          expect(isExisting).toBe(true);
        })
        .call(done);
    });
  });
});
