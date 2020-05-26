import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import { Badge, UncontrolledDropdown, DropdownItem, DropdownMenu, DropdownToggle, Nav, NavItem } from 'reactstrap';
import PropTypes from 'prop-types';

import { AppAsideToggler, AppNavbarBrand, AppSidebarToggler } from '@coreui/react';
import logo from '../../assets/img/brand/logo.svg'
import sygnet from '../../assets/img/brand/sygnet.svg'

const propTypes = {
  children: PropTypes.node,
};

const defaultProps = {};

class DefaultHeader extends Component {
  render() {

    // eslint-disable-next-line
    const { children, ...attributes } = this.props;

    return (
      <React.Fragment>
        <AppSidebarToggler className="d-lg-none" display="md" mobile />
        <AppNavbarBrand
          full={{ src: logo, width: 89, height: 25, alt: 'CoreUI Logo' }}
          minimized={{ src: sygnet, width: 30, height: 30, alt: 'CoreUI Logo' }}
        />
        <AppSidebarToggler className="d-md-down-none" display="lg" />

        <Nav className="d-md-down-none" navbar>
          <NavItem className="px-3">
            <NavLink to="/dashboard" className="nav-link" >Dashboard</NavLink>
          </NavItem>
          <NavItem className="px-3">
            <Link to="/users" className="nav-link">Users</Link>
          </NavItem>
          <NavItem className="px-3">
            <NavLink to="#" className="nav-link">Settings</NavLink>
          </NavItem>
        </Nav>
        <Nav className="ml-auto" navbar>
          <NavItem className="d-md-down-none">
            <NavLink to="#" className="nav-link"><i className="icon-bell"></i><Badge pill color="danger">5</Badge></NavLink>
          </NavItem>
          <NavItem className="d-md-down-none">
            <NavLink to="#" className="nav-link"><i className="icon-list"></i></NavLink>
          </NavItem>
          <NavItem className="d-md-down-none">
            <NavLink to="#" className="nav-link"><i className="icon-location-pin"></i></NavLink>
          </NavItem>
          <UncontrolledDropdown nav direction="down">
            <DropdownToggle nav>
              <img src={'data:image/png;base64,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'} className="img-avatar" alt="admin@bootstrapmaster.com" />
            </DropdownToggle>
            <DropdownMenu right>
              <DropdownItem header tag="div" className="text-center"><strong>Account</strong></DropdownItem>
              <DropdownItem><i className="fa fa-bell-o"></i> Updates<Badge color="info">42</Badge></DropdownItem>
              <DropdownItem><i className="fa fa-envelope-o"></i> Messages<Badge color="success">42</Badge></DropdownItem>
              <DropdownItem><i className="fa fa-tasks"></i> Tasks<Badge color="danger">42</Badge></DropdownItem>
              <DropdownItem><i className="fa fa-comments"></i> Comments<Badge color="warning">42</Badge></DropdownItem>
              <DropdownItem header tag="div" className="text-center"><strong>Settings</strong></DropdownItem>
              <DropdownItem><i className="fa fa-user"></i> Profile</DropdownItem>
              <DropdownItem><i className="fa fa-wrench"></i> Settings</DropdownItem>
              <DropdownItem><i className="fa fa-usd"></i> Payments<Badge color="secondary">42</Badge></DropdownItem>
              <DropdownItem><i className="fa fa-file"></i> Projects<Badge color="primary">42</Badge></DropdownItem>
              <DropdownItem divider />
              <DropdownItem><i className="fa fa-shield"></i> Lock Account</DropdownItem>
              <DropdownItem onClick={e => this.props.onLogout(e)}><i className="fa fa-lock"></i> Logout</DropdownItem>
            </DropdownMenu>
          </UncontrolledDropdown>
        </Nav>
        <AppAsideToggler className="d-md-down-none" />
        {/*<AppAsideToggler className="d-lg-none" mobile />*/}
      </React.Fragment>
    );
  }
}

DefaultHeader.propTypes = propTypes;
DefaultHeader.defaultProps = defaultProps;

export default DefaultHeader;
