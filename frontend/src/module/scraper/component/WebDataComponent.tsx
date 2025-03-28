import React, { useEffect, useState } from 'react';
import { WebSiteDataDto } from '../../../model/scarper';

interface WebDataComponentProps {
  data: WebSiteDataDto;
}

const WebDataComponent = ({ data }: WebDataComponentProps) => {
  const [externalLinks, setExternalLinks] = useState([]);
  const [internalLinks, setInternalLinks] = useState([]);
  const [inaccessibleLinks, setInaccessibleLinks] = useState([]);

  useEffect(() => {
    if (data?.links) {
      setExternalLinks(data?.links?.external);
      setInternalLinks(data?.links?.internal);
      setInaccessibleLinks(data?.links?.inaccessible);
    }
  }, [data]);

  return (
    <div className="container right-sidebar">
      <div className="row mt-5 mb-5">
        <div className="col-sm-12 mb-1">
          <p className="sub-text">SITE URL</p>
          <h5>{data?.url}</h5>
        </div>

        <hr />

        <div className="col-sm-12 mb-1">
          <p className="sub-text">TITLE</p>
          <h5>{data?.title}</h5>
        </div>

        <hr />

        <div className="col-sm-12 mb-1">
          <p className="sub-text">HTML VERSION</p>
          <h5>{data?.html_version}</h5>
        </div>

        <hr />

        <div className="col-sm-12 mb-1">
          <p className="sub-text">Heading count</p>
          <h5>{`H1 (${data?.heading_levels['h1']}), 
          H2 (${data?.heading_levels['h2']}),
          H3 (${data?.heading_levels['h3']}),
          H4 (${data?.heading_levels['h4']}),
          H5 (${data?.heading_levels['h5']}),
          H6 (${data?.heading_levels['h6']})`}</h5>
        </div>

        <hr></hr>

        <div className="col-sm-12 mb-1">
          <p className="sub-text">HAVE LOGIN FORM</p>
          <h5>{data?.have_login_form ? 'YES' : 'NO'}</h5>
        </div>

        <hr />

        <div className="col-sm-12 mb-4">
          <p className="sub-text">External Links - {externalLinks?.length}</p>
          {externalLinks?.map((item) => {
            return (
              <>
                <a className="url">{item}</a>
                <br />
              </>
            );
          })}
        </div>

        <div className="col-sm-12 mb-4">
          <p className="sub-text">Internal Links - {internalLinks?.length}</p>
          {internalLinks?.map((item) => {
            return (
              <>
                <a className="url">{item}</a>
                <br />
              </>
            );
          })}
        </div>

        <div className="col-sm-12 mb-4">
          <p className="sub-text">InaccessibleLinks Links - {inaccessibleLinks?.length}</p>
          {inaccessibleLinks?.map((item) => {
            return (
              <>
                <a className="url">{item}</a>
                <br />
              </>
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default WebDataComponent;
