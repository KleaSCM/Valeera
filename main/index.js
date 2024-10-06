const { app, BrowserWindow } = require('electron');
const path = require('path');

app.disableHardwareAcceleration();

function createWindow() {
    const win = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            preload: path.join(__dirname, 'preload.js'), // Preload script
            contextIsolation: true, // Secure context isolation
            enableRemoteModule: false, // No remote module
            nodeIntegration: false, // Node.js disabled in the renderer
        },
    });

    win.loadURL('http://localhost:3000'); // Load Next.js frontend
}

app.on('ready', createWindow);

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    }
});
