var admin = require("./info.js"),
  user = require("../general/user.js");

accountsPage = {
  title: 'Admin - Accounts'
};

describe('Admin - Accounts', function() {
  admin.init(browser);
  it('navigate to page', function(done) {
      browser
      .element("#sidebar")
      .click("a=Accounts")
      .getTitle(function(err, title) {
          expect(err).toBe(undefined);
          expect(title).toBe(accountsPage.title);
      })
      .call(done);
  });
  it('goes to the page which modifies the account type', function(done){
      browser.
      element('#list-accounts tbody tr:nth-Child(3) td:nth-Child(5)').
      click('a=Modify')
      .pause(1000)
      .call(done);
  });
  it('modifies the type of the account from user to administrator', function(done){
      browser
      .selectByValue('#modify-administrator', 'true')
      .submitForm("#modify")
      .getText(".alert-success").then(function(text){
          expect(text).toContain("The account has been modified!");
      })
      .pause(500)
      .click("a=Go back")
      .call(done);
  });
  it('deletes an account', function(done){
      browser
      .url("/admin")
      .url("/admin/accounts")
      .pause(1000)
      .element('#list-accounts tbody tr:nth-Child(1)')
      .click('a=Delete')
      .pause(1000)
      .keys("U+E00C")
      .call(done);
  });
});
