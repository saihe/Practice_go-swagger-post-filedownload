import React, { Component } from 'react';
import FileDownloadIcon from '@mui/icons-material/FileDownload';
import Button from '@mui/material/Button';
import api from './Api'
import './App.css';

const fileDownload = require('js-file-download');

class App extends Component {
  render() {
    return (
      <div className="App">
        <Button
          variant="contained"
          startIcon={<FileDownloadIcon />}
          onClick={() => {
            api.post(
              '/file',
              {
                id: '0001',
                name: 'download'
              }
            ).then((res) => {
              console.log(res);
              console.log(atob(res.data.data));
              fileDownload(atob(res.data.data), res.data.name);
            });
          }}
        >ダウンロード</Button>
      </div>
    );
  }
}

export default App;
