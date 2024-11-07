package mailers

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendTicket(qrFilePath, email string) error {
	defer os.Remove(qrFilePath)
	qrFileName := qrFilePath[14:] // 14 is the len("./tempTickets/")
	message := gomail.NewMessage()
	message.SetHeader("From", OUR_EMAIL)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Your Maaya Concert Ticket!")

	message.Embed(qrFilePath)
	body := getBody(qrFileName)
	message.SetBody("text/html", body)

	for i := 0; i < 3; i++ {
		// retry up to 3 times
		if err := DIALER.DialAndSend(message); err != nil {
			log.Println("attempt", i+1, "failed to send email: ", err)
			if i == 2 {
				// it failed on the last attempt. Try no more
				return errors.New("failed to send email")
			}
			continue
		}
		break
	}

	log.Println("ticket sent successfully to", email)
	return nil
}

// I know this function is shitty. But I can't think of a better way to do this. And the concert is tomorrow or day after idk
// if I try to load the entire html from a file, and replace the img tag's src through fmt.Sprintf on the
// entire html, it replaces it in unintended places. Since the html contians css which contains % symbols
func getBody(qrFileName string) string {
	body := `
<!DOCTYPE html>
<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="en">

<head>
  <title></title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]--><!--[if !mso]><!-->
  <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@100;200;300;400;500;600;700;800;900" rel="stylesheet"
    type="text/css"><!--<![endif]-->
  <style>
    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      padding: 0;
    }

    a[x-apple-data-detectors] {
      color: inherit !important;
      text-decoration: inherit !important;
    }

    #MessageViewBody a {
      color: inherit;
      text-decoration: none;
    }

    p {
      line-height: inherit
    }

    .desktop_hide,
    .desktop_hide table {
      mso-hide: all;
      display: none;
      max-height: 0px;
      overflow: hidden;
    }

    .image_block img+div {
      display: none;
    }

    sup,
    sub {
      font-size: 75%;
      line-height: 0;
    }

    @media (max-width:520px) {
      .desktop_hide table.icons-inner {
        display: inline-block !important;
      }

      .icons-inner {
        text-align: center;
      }

      .icons-inner td {
        margin: 0 auto;
      }

      .mobile_hide {
        display: none;
      }

      .row-content {
        width: 100% !important;
      }

      .stack .column {
        width: 100%;
        display: block;
      }

      .mobile_hide {
        min-height: 0;
        max-height: 0;
        max-width: 0;
        overflow: hidden;
        font-size: 0px;
      }

      .desktop_hide,
      .desktop_hide table {
        display: table !important;
        max-height: none !important;
      }
    }
  </style>
  <!--[if mso ]><style>sup, sub { font-size: 100% !important; } sup { mso-text-raise:10% } sub { mso-text-raise:-10% }</style> <![endif]-->
</head>

<body class="body"
  style="background-color: #0a123a; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;">
  <table class="nl-container" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation"
    style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #0a123a;">
    <tbody>
      <tr>
        <td>
          <table class="row row-1" align="center" width="100%" border="0" cellpadding="0" cellspacing="0"
            role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
            <tbody>
              <tr>
                <td>
                  <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0"
                    role="presentation"
                    style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px; margin: 0 auto;"
                    width="500">
                    <tbody>
                      <tr>
                        <td class="column column-1" width="100%"
                          style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; padding-bottom: 5px; padding-top: 5px; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;">
                          <table class="heading_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0"
                            role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
                            <tr>
                              <td class="pad">
                                <h1
                                  style="margin: 0; color: #f6f668; direction: ltr; font-family: 'Roboto', Tahoma, Verdana, Segoe, sans-serif; font-size: 38px; font-weight: 700; letter-spacing: normal; line-height: 120%; text-align: left; margin-top: 0; margin-bottom: 0; mso-line-height-alt: 45.6px;">
                                  <span class="tinyMce-placeholder" style="word-break: break-word;">Your Maaya Concert
                                    Ticket<br></span>
                                </h1>
                              </td>
                            </tr>
                          </table>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </td>
              </tr>
            </tbody>
          </table>
          <table class="row row-2" align="center" width="100%" border="0" cellpadding="0" cellspacing="0"
            role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
            <tbody>
              <tr>
                <td>
                  <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0"
                    role="presentation"
                    style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; border-radius: 0; color: #000000; width: 500px; margin: 0 auto;"
                    width="500">
                    <tbody>
                      <tr>
                        <td class="column column-1" width="100%"
                          style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; padding-bottom: 5px; padding-top: 5px; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;">
`

	body = body + fmt.Sprintf(`<img src="cid:%s" width="512px"></img>`, qrFileName)

	body = body + `
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </td>
              </tr>
            </tbody>
          </table>
          <table class="row row-3" align="center" width="100%" border="0" cellpadding="0" cellspacing="0"
            role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
            <tbody>
              <tr>
                <td>
                  <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0"
                    role="presentation"
                    style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; border-radius: 0; color: #000000; width: 500px; margin: 0 auto;"
                    width="500">
                    <tbody>
                      <tr>
                        <td class="column column-1" width="100%"
                          style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; padding-bottom: 5px; padding-top: 5px; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;">
                          <table class="paragraph_block block-1" width="100%" border="0" cellpadding="10"
                            cellspacing="0" role="presentation"
                            style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
                            <tr>
                              <td class="pad">
                                <div
                                  style="color:#fffffc;direction:ltr;font-family:Arial, 'Helvetica Neue', Helvetica, sans-serif;font-size:16px;font-weight:400;letter-spacing:0px;line-height:120%;text-align:left;mso-line-height-alt:19.2px;">
                                  <p style="margin: 0; margin-bottom: 16px;"><strong>Please keep this ready at the time
                                      of entry to the concert. You will not be allowed entry without this
                                      ticket</strong></p>
                                  <p style="margin: 0; margin-bottom: 16px;">&nbsp;</p>
                                  <p style="margin: 0; margin-bottom: 16px;"><strong>Rules:</strong></p>
                                  <p style="margin: 0; margin-bottom: 16px;">1. Admission to the event is only permitted
                                    with a valid digital ticket </p>
                                  <p style="margin: 0; margin-bottom: 16px;">2. Each digital ticket grants entry to one
                                    person only.</p>
                                  <p style="margin: 0; margin-bottom: 16px;">3. For the safety of all attendees, event
                                    organizers reserve the right to refuse admission and conduct security searches as
                                    they deem appropriate.</p>
                                  <p style="margin: 0; margin-bottom: 16px;">4. Digital tickets are non-transferable,
                                    non-exchangeable, and non-refundable after purchase for any reason.</p>
                                  <p style="margin: 0; margin-bottom: 16px;">5. Reselling digital tickets for commercial
                                    gain is strictly prohibited.</p>
                                  <p style="margin: 0; margin-bottom: 16px;">6. Professional cameras and recording
                                    equipment are not permitted inside the event premises.</p>
                                  <p style="margin: 0; margin-bottom: 16px;">7. Re-entry is not allowed; participants
                                    will not be permitted to return to the concert area if they leave the campus
                                    premises.</p>
                                  <p style="margin: 0; margin-bottom: 16px;">8. The use of tobacco products, alcohol, or
                                    any form of pre-gaming is strictly prohibited.</p>
                                  <p style="margin: 0;">9. Event organizers will not be responsible for any loss or
                                    damage to personal belongings.</p>
                                </div>
                              </td>
                            </tr>
                          </table>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </td>
              </tr>
            </tbody>
          </table>
        </td>
      </tr>
    </tbody>
  </table><!-- End -->
</body>

</html>
`

	return body
}
