
describe('Shortcut Keys', function() {
  it('ALT + A', function(done) {
      browser
      .url("/")
      .click("#wisply-body")
      .pause(1000)
      .keys(["U+E00AA"])
      .pause(1500)
      .isExisting('#at-collapse').then(function(isExisting){
          expect(isExisting).toBe(true);
      })
      .call(done);
  });

});
