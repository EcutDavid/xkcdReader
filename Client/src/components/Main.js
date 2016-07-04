import React from 'react'
import request from 'superagent'
import 'normalize.css/normalize.css'

import 'styles/app.scss'

class AppComponent extends React.Component {
  componentDidMount() {
    request
    .get('http://xkcd.com/614/info.0.json')
    .end(function(err, res){
      console.log(err)
      console.log(res)
    });
  }

  render() {
    return (
      <div className="index">
        <button className="button">Hello human</button>
      </div>
    )
  }
}

export default AppComponent
