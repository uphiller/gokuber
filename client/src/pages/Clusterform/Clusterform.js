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

class Clusterform extends Component {
  constructor(props) {
    super(props);

    this.state = {
      name:'',
      type: '',
      quntity:''
    };
    this.handleSave = this.handleSave.bind(this);
    this.handleNameChange = this.handleNameChange.bind(this);
    this.handleTypeChange = this.handleTypeChange.bind(this);
    this.handleQuntityChange = this.handleQuntityChange.bind(this);
  }

  handleNameChange(e){
    e.preventDefault();
    this.setState({name:e.target.value})
  }

  handleTypeChange(e){
    e.preventDefault();
    this.setState({type:e.target.value})
  }

  handleQuntityChange(e){
    e.preventDefault();
    this.setState({quntity:e.target.value})
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
    if(!this.state.quntity){
      NotificationManager.warning('Error message', 'Quntity is Required value!', 5000, () => {
        alert('callback');
      });
      return;
    }
    const response = await service.setCluster(this.state);

    NotificationManager.success('Save is success!!', 'success');
    this.props.history.push('/clusters');
  }

  render() {
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" sm="12">
            <Card>
              <CardHeader>
                <strong>Cluster</strong>
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
                  <Col xs="4">
                    <FormGroup>
                      <Label htmlFor="ccmonth">Node Quantity</Label>
                      <Input type="select" name="ccmonth" id="ccmonth" value={this.state.quntity} onChange={this.handleQuntityChange}>
                        <option value="">Quantity select</option>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                        <option value="4">4</option>
                        <option value="5">5</option>
                      </Input>
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

export default Clusterform;
