
describe('Shortcut Keys', function() {
  it('ALT + A', function(done) {
      browser
      .url("/")
      .keys(['ALT', 'A'])
      .pause(1500)
      .isExisting('#at-collapse').then(function(isExisting){
          expect(isExisting).toBe(true);
      })
      .call(done);
  });

});
