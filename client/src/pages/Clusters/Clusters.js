import React, { Component } from 'react';
import { Badge, Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Row, Table } from 'reactstrap';
import * as service from "../../rest";
import axios from "axios";

class Tables extends Component {

  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      clusters: []
    };
  }

  async componentDidMount() {
    let response = await service.getGcpClusters(this.state);
    this.setState({
      clusters: response.data.clusters
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
                    <th>Status</th>
                  </tr>
                  </thead>
                  <tbody>

                    {this.state.clusters.map((value, index) => {
                      return  (<tr>
                                  <td>{value.Name}</td>
                                  <td>{value.CreatedAt}</td>
                                  <td>Member</td>
                                  <td>
                                    <Badge color="success">{value.Status}</Badge>
                                  </td>
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
              </CardBody>
            </Card>
          </Col>
        </Row>
      </div>

    );
  }
}

export default Tables;
