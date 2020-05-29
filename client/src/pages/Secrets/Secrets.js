import React, { Component } from 'react';
import {
  Badge,
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Pagination,
  PaginationItem,
  PaginationLink,
  Row,
  Table
} from 'reactstrap';
import * as service from "../../rest";

class Secrets extends Component {

  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      secrets: []
    };
  }

  async componentDidMount() {
    let response = await service.getSecrets();
    this.setState({
      secrets: response.data.secrets
    });
  }

  render() {
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" lg="12">
            <Card>
              <CardHeader>
                <i className="fa fa-align-justify"></i> Clusters
              </CardHeader>
              <CardBody>
                <Table responsive>
                  <thead>
                  <tr>
                    <th>Name</th>
                    <th>Date registered</th>
                    <th>Create User</th>
                  </tr>
                  </thead>
                  <tbody>

                  {this.state.secrets.map((value, index) => {
                    return  (<tr>
                      <td>{value.Name}</td>
                      <td>{value.CreatedAt}</td>
                      <td>Member</td>
                    </tr>)
                  })}
                  </tbody>
                </Table>
                <Pagination>
                  <PaginationItem>
                    <PaginationLink previous tag="button"></PaginationLink>
                  </PaginationItem>
                  <PaginationItem active>
                    <PaginationLink tag="button">1</PaginationLink>
                  </PaginationItem>
                  <PaginationItem>
                    <PaginationLink tag="button">2</PaginationLink>
                  </PaginationItem>
                  <PaginationItem>
                    <PaginationLink tag="button">3</PaginationLink>
                  </PaginationItem>
                  <PaginationItem>
                    <PaginationLink tag="button">4</PaginationLink>
                  </PaginationItem>
                  <PaginationItem>
                    <PaginationLink next tag="button"></PaginationLink>
                  </PaginationItem>
                </Pagination>
                <Col sm xs="12" className="text-center mt-3">
                  <Button color="primary" onClick={() => this.props.history.push('/secretform')}>
                    <i className="fa fa-lightbulb-o"></i>&nbsp;Create Secret
                  </Button>
                </Col>
              </CardBody>
            </Card>
          </Col>
        </Row>
      </div>

    );
  }
}

export default Secrets;
