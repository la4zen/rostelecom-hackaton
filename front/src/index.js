import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter } from 'react-router-dom';
import App from './router/App';
import Header from './router/Header'

ReactDOM.render(
  <BrowserRouter>
    <Header/>
    <App/>
  </BrowserRouter>,
  document.getElementById('root'));

reportWebVitals();
