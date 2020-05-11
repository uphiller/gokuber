import React, { Component } from 'react';
import logo from './logo.svg';
// import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import axios from 'axios'

class App extends Component {
    /* id password state값 으로 정의 */
    state = {
        subscriptionId: '',
        clientId: '',
        clientSecret: '',
        tenantId:''
    }
    /* input value 변경 ==> onChange */
    appChange = (e) => {
        this.setState({
            [e.target.name]: e.target.value
        });
    }
    appClick = () => {
        axios.post('http://localhost:5000/v1/info', {
            subscriptionId: this.state.subscriptionId,
            clientId: this.state.clientId,
            clientSecret:this.state.clientSecret,
            tenantId:this.state.tenantId
        })
        .then( response => { console.log(response) } )
        .catch( response => { console.log(response) } );
    }
    appKeyPress = (e) => {
        if (e.key === 'Enter') {
            this.appClick();
        }
    }

    render() {
        const { subscriptionId, clientId, clientSecret,  tenantId} = this.state;
        const { appChange, appClick, appKeyPress } = this;
        return (
            <div className="App">
                <header className="App-header">
                    <input type="text" name="subscriptionId" placeholder="아이디" value={subscriptionId} onChange={appChange}/>
                    <input type="text" name="clientId" placeholder="아이디" value={clientId} onChange={appChange}/>
                    <input type="text" name="clientSecret" placeholder="아이디" value={clientSecret} onChange={appChange}/>
                    <input type="text" name="tenantId" placeholder="아이디" value={tenantId} onChange={appChange}/>

                    <Button variant="primary" onClick={appClick}>Primary</Button>
                </header>
            </div>
        );
    }
}

export default App;
