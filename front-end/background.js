chrome.runtime.onInstalled.addListener(function () {
  chrome.storage.sync.set({
    color: '#3aa757'
  }, function () {
    console.log("The color is green.");
  });

  chrome.declarativeContent.onPageChanged.removeRules(undefined, function () {
    chrome.declarativeContent.onPageChanged.addRules(Rules());
  });
});

function Rules() {
  return [{
      conditions: [new chrome.declarativeContent.PageStateMatcher({
        pageUrl: {
          hostContains: 'amazon'
        }
      })],
      actions: [new chrome.declarativeContent.ShowPageAction()]
    },
    {
      conditions: [new chrome.declarativeContent.PageStateMatcher({
        pageUrl: {
          hostContains: 'flipkart'
        }
      })],
      actions: [new chrome.declarativeContent.ShowPageAction()]
    }
  ];
}