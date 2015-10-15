

var AdminAccount = {
    name : "Cristian Sima",
    email: "cristian.sima93@yahoo.com",
    password: "password"
};

var UserAccount = {
  name : "User Test",
  email: "test@test.com",
  password: "passwordTest"
};


function loginAccount(account) {
  return function(browser) {
    browser
      .open('/auth/login')
      .type('login-email', account.email)
      .type('login-password', account.password)
      .clickAndWait('login-submit-button')
      .assertTextPresent("Hi, " + account.name);
  };
}

function logoutAccount() {
  return function(browser) {
    browser
      .clickAndWait('menu-logout-button');
  };
}


function success(err) {
    console.log('[Success] The test passed! :) ');
    if(err) throw err;
}

function TestLogin() {
  return function(browser) {

  };
}

exports.Name = "Authentification system";
exports.Start = function(browser) {
  browser
    .chain
    .session()
    .and(loginAccount(AdminAccount))
    .and(logoutAccount())
    .testComplete()
    .end(function(err){
      browser.testComplete(success);
    });
};
