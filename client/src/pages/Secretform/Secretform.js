import React, { Component } from 'react';
import {
  Badge,
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Col,
  Collapse,
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
  Fade,
  Form,
  FormGroup,
  FormText,
  FormFeedback,
  Input,
  InputGroup,
  InputGroupAddon,
  InputGroupButtonDropdown,
  InputGroupText,
  Label,
  Row,
} from 'reactstrap';
import * as service from "../../rest";
import {NotificationManager} from "react-notifications";

class Secretform extends Component {
  constructor(props) {
    super(props);

    this.state = {
      name:'',
      access_id: '',
      secret_key:'',
      type:''
    };
    this.handleSave = this.handleSave.bind(this);
    this.handleNameChange = this.handleNameChange.bind(this);
    this.handleAccessIdChange = this.handleAccessIdChange.bind(this);
    this.handleSecretKeyChange = this.handleSecretKeyChange.bind(this);
    this.handleTypeChange = this.handleTypeChange.bind(this);
  }

  handleNameChange(e){
    e.preventDefault();
    this.setState({name:e.target.value})
  }

  handleAccessIdChange(e){
    e.preventDefault();
    this.setState({access_id:e.target.value})
  }

  handleSecretKeyChange(e){
    e.preventDefault();
    this.setState({secret_key:e.target.value})
  }

  handleTypeChange(e){
    e.preventDefault();
    this.setState({type:e.target.value})
  }

  handleSave = async () => {
    if(!this.state.name){
      NotificationManager.warning('Error message', 'Name is Required value!', 5000, () => {
        alert('callback');
      });
      return;
    }
    if(!this.state.type){
      NotificationManager.warning('Error message', 'Type is Required value!', 5000, () => {
        alert('callback');
      });
      return;
    }
    if(!this.state.access_id){
      NotificationManager.warning('Error message', 'Access Id is Required value!', 5000, () => {
        alert('callback');
      });
      return;
    }
    if(!this.state.secret_key){
      NotificationManager.warning('Error message', 'Secret Key is Required value!', 5000, () => {
        alert('callback');
      });
      return;
    }
    const response = await service.setSecret(this.state);

    NotificationManager.success('Save is success!!', 'success');
    this.props.history.push('/secrets');
  }

  render() {
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" sm="12">
            <Card>
              <CardHeader>
                <strong>Secret</strong>
                <small> Form</small>
              </CardHeader>
              <CardBody>
                <Row>
                  <Col xs="12">
                    <FormGroup>
                      <Label htmlFor="name">Name</Label>
                      <Input type="text" id="name" placeholder="Enter cluster name" required value={this.state.name}  onChange={this.handleNameChange}/>
                    </FormGroup>
                  </Col>
                </Row>
                <Row>
                  <Col xs="12">
                    <FormGroup>
                      <Label htmlFor="name">Type</Label>
                      <Input type="select" name="select" id="select" value={this.state.type} onChange={this.handleTypeChange}>
                        <option value="">Cloud select</option>
                        <option value="aws">AWS</option>
                        <option value="azur">AZUR</option>
                        <option value="gcp">GCP</option>
                      </Input>
                    </FormGroup>
                  </Col>
                </Row>
                <Row>
                  <Col xs="12">
                    <FormGroup>
                      <Label htmlFor="name">Access Id</Label>
                      <Input type="text" placeholder="Enter cluster name" required value={this.state.access_id}  onChange={this.handleAccessIdChange}/>
                    </FormGroup>
                  </Col>
                </Row>
                <Row>
                  <Col xs="12">
                    <FormGroup>
                      <Label htmlFor="name">Secret Key</Label>
                      <Input type="text" placeholder="Enter cluster name" required value={this.state.secret_key}  onChange={this.handleSecretKeyChange}/>
                    </FormGroup>
                  </Col>
                </Row>
              </CardBody>
              <CardFooter>
                <Button type="submit" size="sm" color="primary" onClick={this.handleSave} ><i className="fa fa-dot-circle-o"></i> Submit</Button>
              </CardFooter>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

export default Secretform;
