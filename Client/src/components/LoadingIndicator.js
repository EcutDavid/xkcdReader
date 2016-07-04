import React from 'react'

export default class LoadingIndicator extends React.Component {
  constructor() {
    super()
    this.state = { MS100Counter: 0}
  }

  componentWillMount() {
    // setInterval(() => {
    //   let { MS100Counter } = this.state
    //   MS100Counter++
    //   this.setState({ MS100Counter })
    // }, 100)
  }

  render() {
    // const { MS100Counter } = this.state
    return (
      <div>
        <p className='indicator'>Loading comics</p>
      </div>
    )
  }
}
