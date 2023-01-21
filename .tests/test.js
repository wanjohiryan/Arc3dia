(async () => {
    const puppeteer = require('puppeteer');
    const GIFEncoder = require('gif-encoder');
    const getPixels = require('get-pixels');
    
    const WIDTH = 1920
    const HEIGHT = 1080
    
    const encoder = new GIFEncoder(WIDTH, HEIGHT);
    const fs = require('fs');
    const workDir = './temp/';
    let file = require('fs').createWriteStream('test.gif');
    const browser = await puppeteer.launch({
      headless: true,
      args: [
        '--origin-to-force-quic-on=localhost:8080'
      ]
     });
    const page = await browser.newPage();
  
    if (!fs.existsSync(workDir)) {
      fs.mkdirSync(workDir);
    };
  
    // Setup gif encoder parameters
    encoder.setFrameRate(60); //@60fps
    encoder.pipe(file);
    encoder.setQuality(40);
    encoder.setDelay(500);
    encoder.writeHeader();
    encoder.setRepeat(0);
  
    // Helper functions declaration
    function addToGif(images, counter = 0) {
      getPixels(images[counter], function (err, pixels) {
  
        encoder.addFrame(pixels.data);
        encoder.read();
        if (counter === images.length - 1) {
          encoder.finish();
          cleanUp(images, function (err) {
            if (err) {
              console.log(err);
            } else {
              fs.rmdirSync(workDir);
              console.log('Gif created!');
              process.exit(0);
            }
          });
  
        } else {
          addToGif(images, ++counter);
        }
      });
    };
  
    function cleanUp(listOfPNGs, callback) {
      let i = listOfPNGs.length;
      listOfPNGs.forEach(function (filepath) {
        fs.unlink(filepath, function (err) {
          i--;
          if (err) {
            callback(err);
            return;
          } else if (i <= 0) {
            callback(null);
          }
        });
      });
    };
  
    async function scrollPage() {
      await page.evaluate(async () => {
        window.scrollBy(0, 10);
      });
    }
  
    await page.setViewport({ width: WIDTH, height: HEIGHT });
    await page.goto('http://localhost:8080', { timeout: 0});
  
    for (let i = 0; i < 60; i++) {
      await page.screenshot({ path: workDir + i + ".png" });
      //FIXME: don't scroll page, interact with it directly
      await scrollPage();
    }
  
    await browser.close();
  
    let listOfPNGs = fs.readdirSync(workDir)
      .map(a => a.substr(0, a.length - 4) + '')
      .sort(function (a, b) { return a - b })
      .map(a => workDir + a.substr(0, a.length) + '.png');
  
    addToGif(listOfPNGs);
  })();
