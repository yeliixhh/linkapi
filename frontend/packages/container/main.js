const { app, BrowserWindow } = require('electron')

const createWindow = () => {

  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      webSecurity: false,
    }
  })

  win.loadURL('http://localhost:5173/')

  win.webContents.openDevTools()

}

app.whenReady().then(() => {
  createWindow()
})
