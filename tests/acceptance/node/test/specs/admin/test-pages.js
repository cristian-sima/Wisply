var admin = require("./info.js"),
  user = require("../general/user.js");

accountsPage = {
  title: 'Admin - Accounts'
};

describe('Admin - Test pages', function() {
  admin.init(browser);
  it('navigate to institutions', function(done) {
      browser
      .element("#sidebar")
      .click("a=Institutions")
      .getTitle(function(err, title) {
          expect(err).toBe(undefined);
          expect(title).toBe(accountsPage.title);
      })
      .call(done);
  });
});
