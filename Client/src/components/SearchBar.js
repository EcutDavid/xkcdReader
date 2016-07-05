import React from 'react'

import 'styles/searchBar.scss'

export default class SearchBar extends React.Component {
  constructor() {
    super()
    this.state = { loading: 0}
  }
  render() {
    const { loading } = this.state
    return (
      <div className='SearchBar row small-11 small-centered'>
        <input
          className='small-9 columns'
          type='text'
          placeholder='Part of the comic title'
        />
        <button
          className='button'
        >
          <i className={`fa ${loading ? 'fa-spinner fa-spin' : 'fa-search'}`}/>
        </button>
      </div>
    )
  }
}
